package main

import (
	"archive/zip"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Server struct {
	db        *sql.DB
	jobs      map[string]*PendingJob
	mu        sync.RWMutex
	uploadDir string
	dbPath    string
}

func NewServer(db *sql.DB, uploadDir, dbPath string) *Server {
	return &Server{
		db:        db,
		jobs:      make(map[string]*PendingJob),
		uploadDir: uploadDir,
		dbPath:    dbPath,
	}
}

func (s *Server) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/upload", s.handleUpload)
	mux.HandleFunc("GET /api/uploads", s.handleListUploads)
	mux.HandleFunc("DELETE /api/uploads/{id}", s.handleDeleteUpload)
	mux.HandleFunc("GET /api/uploads/{id}/download", s.handleUploadDownload)
	mux.HandleFunc("POST /api/fetch", s.handleFetch)
	mux.HandleFunc("POST /api/cancel", s.handleCancel)
	mux.HandleFunc("GET /api/jobs", s.handleListJobs)
	mux.HandleFunc("GET /api/jobs/{id}", s.handleGetJob)
	mux.HandleFunc("GET /api/jobs/{id}/sse", s.handleJobSSE)
	mux.HandleFunc("GET /api/jobs/{id}/download", s.handleDownload)
	mux.HandleFunc("GET /api/settings", s.handleGetSettings)
	mux.HandleFunc("POST /api/settings", s.handleUpdateSettings)
	mux.HandleFunc("GET /api/export", s.handleExport)
}

func jsonResponse(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if v != nil {
		json.NewEncoder(w).Encode(v)
	}
}

func jsonError(w http.ResponseWriter, status int, msg string) {
	jsonResponse(w, status, map[string]string{"error": msg})
}

type uploadResponse struct {
	FileID   string `json:"file_id"`
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
}

func (s *Server) handleUpload(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(50 << 20); err != nil {
		jsonError(w, 400, "file too large or invalid form")
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		jsonError(w, 400, "no file provided")
		return
	}
	defer file.Close()

	ext := filepath.Ext(header.Filename)
	if strings.ToLower(ext) != ".xlsx" {
		jsonError(w, 400, "only .xlsx files accepted")
		return
	}

	if err := os.MkdirAll(s.uploadDir, 0755); err != nil {
		jsonError(w, 500, "server error")
		return
	}

	fileID := uuid.New().String()
	diskPath := filepath.Join(s.uploadDir, fileID+".xlsx")

	dst, err := os.Create(diskPath)
	if err != nil {
		jsonError(w, 500, "failed to save file")
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		jsonError(w, 500, "failed to write file")
		return
	}

	if err := insertUpload(s.db, fileID, header.Filename, diskPath); err != nil {
		jsonError(w, 500, "failed to track upload")
		return
	}

	activeFileID, _ := getSetting(s.db, "active_file_id")
	if activeFileID == "" {
		_ = setSetting(s.db, "active_file_id", fileID)
	}

	jsonResponse(w, 200, uploadResponse{
		FileID:   fileID,
		FileName: header.Filename,
		FilePath: diskPath,
	})
}

type fetchRequest struct {
	FileID                 *string  `json:"file_id,omitempty"`
	FilePath               *string  `json:"file_path,omitempty"`
	UploadID               *string  `json:"upload_id,omitempty"`
	StartNumber            int      `json:"startNumber"`
	EndNumber              int      `json:"endNumber"`
	SheetName              string   `json:"sheetName"`
	SheetNames             []string `json:"sheetNames,omitempty"`
	ArtikelnummerCol       string   `json:"artikelnummerCol"`
	PDFLinkCol             string   `json:"pdfLinkCol"`
	DocumentType           string   `json:"documentType"`
	UseDocumentTypeColumn  bool     `json:"use_document_type_column"`
	DocumentTypeColumn     string   `json:"document_type_column"`
	Company                string   `json:"company"`
	UpdateDate             bool     `json:"update_date"`
	DateColumn             string   `json:"date_column"`
	Concurrency            int      `json:"concurrency"`
}

type fetchResponse struct {
	JobID string `json:"job_id"`
}

func (s *Server) handleFetch(w http.ResponseWriter, r *http.Request) {
	var req fetchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonError(w, 400, "invalid JSON body")
		return
	}

	var filePath, fileName string
	if req.UploadID != nil && *req.UploadID != "" {
		upload, err := getUpload(s.db, *req.UploadID)
		if err != nil {
			jsonError(w, 400, "upload not found")
			return
		}
		if _, err := os.Stat(upload.FilePath); os.IsNotExist(err) {
			jsonError(w, 400, "uploaded file not found on disk")
			return
		}
		filePath = upload.FilePath
		fileName = upload.OriginalName
	} else if req.FileID != nil && *req.FileID != "" {
		p := filepath.Join(s.uploadDir, *req.FileID+".xlsx")
		if _, err := os.Stat(p); os.IsNotExist(err) {
			jsonError(w, 400, "uploaded file not found")
			return
		}
		filePath = p
		fileName = *req.FileID + ".xlsx"
	} else if req.FilePath != nil && *req.FilePath != "" {
		if _, err := os.Stat(*req.FilePath); os.IsNotExist(err) {
			jsonError(w, 400, "server file path not found")
			return
		}
		filePath = *req.FilePath
		fileName = filepath.Base(*req.FilePath)
	} else {
		jsonError(w, 400, "file_id, file_path or upload_id required")
		return
	}

	params := FetchParams{
		File:                  filePath,
		StartNumber:           req.StartNumber,
		EndNumber:             req.EndNumber,
		SheetName:             req.SheetName,
		SheetNames:            req.SheetNames,
		ArtikelnummerCol:      req.ArtikelnummerCol,
		PDFLinkCol:            req.PDFLinkCol,
		DocumentType:          req.DocumentType,
		UseDocumentTypeColumn: req.UseDocumentTypeColumn,
		DocumentTypeColumn:    req.DocumentTypeColumn,
		Company:               req.Company,
		UpdateDate:            req.UpdateDate,
		DateColumn:            req.DateColumn,
		Concurrency:           req.Concurrency,
	}

	if err := validateParams(params); err != nil {
		jsonError(w, 400, err.Error())
		return
	}

	jobID := uuid.New().String()
	ctx, cancel := context.WithCancel(context.Background())

	paramsJSON, _ := json.Marshal(req)
	_ = setSetting(s.db, "last_fetch_params", string(paramsJSON))

	_, err := insertJob(s.db, jobID, fileName, filePath, params, paramsJSON)
	if err != nil {
		cancel()
		jsonError(w, 500, "failed to create job")
		return
	}

	broadcaster := newJobBroadcaster()
	job := &PendingJob{
		ID:          jobID,
		Params:      params,
		FilePath:    filePath,
		FileName:    fileName,
		Ctx:         ctx,
		Cancel:      cancel,
		Broadcaster: broadcaster,
	}

	s.mu.Lock()
	s.jobs[jobID] = job
	s.mu.Unlock()

	go s.runJob(job)

	jsonResponse(w, 200, fetchResponse{JobID: jobID})
}

type cancelRequest struct {
	JobID string `json:"job_id"`
}

func (s *Server) handleCancel(w http.ResponseWriter, r *http.Request) {
	var req cancelRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonError(w, 400, "invalid JSON")
		return
	}

	s.mu.RLock()
	job, ok := s.jobs[req.JobID]
	s.mu.RUnlock()

	if !ok {
		jsonError(w, 404, "job not found or already finished")
		return
	}

	job.Cancel()
	jsonResponse(w, 200, map[string]string{"status": "cancelling"})
}

func (s *Server) handleListJobs(w http.ResponseWriter, r *http.Request) {
	jobs, err := listJobs(s.db)
	if err != nil {
		jsonError(w, 500, "database error")
		return
	}
	if jobs == nil {
		jobs = []jobRow{}
	}
	jsonResponse(w, 200, jobs)
}

func (s *Server) handleGetJob(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	job, err := getJobRow(s.db, id)
	if err != nil {
		jsonError(w, 404, "job not found")
		return
	}
	jsonResponse(w, 200, job)
}

func (s *Server) handleJobSSE(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	s.mu.RLock()
	job, ok := s.jobs[id]
	s.mu.RUnlock()

	if !ok {
		if _, err := getJobRow(s.db, id); err == nil {
			w.Header().Set("Content-Type", "text/event-stream")
			w.Header().Set("Cache-Control", "no-cache")
			fmt.Fprintf(w, "data: {\"status\":\"completed\"}\n\n")
			return
		}
		jsonError(w, 404, "job not found")
		return
	}

	flusher, ok := w.(http.Flusher)
	if !ok {
		jsonError(w, 500, "streaming not supported")
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	ch, unsub := job.Broadcaster.subscribe()
	defer unsub()

	ctx := r.Context()
	for {
		select {
		case <-ctx.Done():
			return
		case update, ok := <-ch:
			if !ok {
				return
			}
			b, _ := json.Marshal(update)
			fmt.Fprintf(w, "data: %s\n\n", b)
			flusher.Flush()
			if update.Status == "completed" || update.Status == "cancelled" || update.Status == "error" {
				return
			}
		}
	}
}

func (s *Server) handleDownload(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	s.mu.RLock()
	job, ok := s.jobs[id]
	s.mu.RUnlock()

	var filePath, fileName string
	if ok {
		filePath = job.FilePath
		fileName = job.FileName
	} else {
		row, err := getJobRow(s.db, id)
		if err != nil {
			jsonError(w, 404, "job not found")
			return
		}
		filePath = row.FilePath
		fileName = row.FileName
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		jsonError(w, 404, "file not found on disk")
		return
	}

	downloadName := strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "_processed" + filepath.Ext(fileName)
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, downloadName))
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	http.ServeFile(w, r, filePath)
}

func (s *Server) handleListUploads(w http.ResponseWriter, r *http.Request) {
	uploads, err := listUploads(s.db)
	if err != nil {
		jsonError(w, 500, "database error")
		return
	}
	jsonResponse(w, 200, uploads)
}

func (s *Server) handleDeleteUpload(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	upload, err := getUpload(s.db, id)
	if err != nil {
		jsonError(w, 404, "upload not found")
		return
	}
	if err := os.Remove(upload.FilePath); err != nil {
		jsonError(w, 500, "failed to delete file")
		return
	}
	if err := deleteUpload(s.db, id); err != nil {
		jsonError(w, 500, "failed to delete upload record")
		return
	}
	activeFileID, _ := getSetting(s.db, "active_file_id")
	if activeFileID == id {
		_ = setSetting(s.db, "active_file_id", "")
	}
	jsonResponse(w, 200, map[string]string{"status": "deleted"})
}

func (s *Server) handleUploadDownload(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	upload, err := getUpload(s.db, id)
	if err != nil {
		jsonError(w, 404, "upload not found")
		return
	}
	if _, err := os.Stat(upload.FilePath); os.IsNotExist(err) {
		jsonError(w, 404, "file not found on disk")
		return
	}
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, upload.OriginalName))
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	http.ServeFile(w, r, upload.FilePath)
}

func (s *Server) handleGetSettings(w http.ResponseWriter, r *http.Request) {
	settings, err := getAllSettings(s.db)
	if err != nil {
		jsonError(w, 500, "database error")
		return
	}
	jsonResponse(w, 200, settings)
}

func (s *Server) handleUpdateSettings(w http.ResponseWriter, r *http.Request) {
	var updates map[string]string
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		jsonError(w, 400, "invalid JSON")
		return
	}
	for k, v := range updates {
		if err := setSetting(s.db, k, v); err != nil {
			jsonError(w, 500, "failed to save setting")
			return
		}
	}
	jsonResponse(w, 200, map[string]string{"status": "ok"})
}

func (s *Server) handleExport(w http.ResponseWriter, r *http.Request) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	timestamp := time.Now().Format("2006-01-02_15-04-05")
	filename := fmt.Sprintf("tremco-export_%s.zip", timestamp)

	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, filename))
	w.Header().Set("Content-Type", "application/zip")

	zw := zip.NewWriter(w)
	defer zw.Close()

	s.addDirToZip(zw, s.uploadDir, "uploads")

	dbFile, err := os.Open(s.dbPath)
	if err == nil {
		defer dbFile.Close()
		dbInfo, err := dbFile.Stat()
		if err == nil {
			header, err := zip.FileInfoHeader(dbInfo)
			if err == nil {
				header.Name = "tremco.db"
				header.Method = zip.Deflate
				w, err := zw.CreateHeader(header)
				if err == nil {
					io.Copy(w, dbFile)
				}
			}
		}
	}
}

func (s *Server) addDirToZip(zw *zip.Writer, dir, prefix string) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return
	}
	for _, entry := range entries {
		path := filepath.Join(dir, entry.Name())
		if entry.IsDir() {
			s.addDirToZip(zw, path, prefix+"/"+entry.Name())
			continue
		}
		file, err := os.Open(path)
		if err != nil {
			continue
		}
		info, err := file.Stat()
		if err != nil {
			file.Close()
			continue
		}
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			file.Close()
			continue
		}
		header.Name = prefix + "/" + entry.Name()
		header.Method = zip.Deflate
		w, err := zw.CreateHeader(header)
		if err != nil {
			file.Close()
			continue
		}
		io.Copy(w, file)
		file.Close()
	}
}

func (s *Server) runJob(job *PendingJob) {
	defer func() {
		s.mu.Lock()
		delete(s.jobs, job.ID)
		s.mu.Unlock()
		job.Cancel()
	}()

	var processed, total int
	report := func(p, t int) {
		processed = p
		total = t
		progress := 0.0
		if t > 0 {
			progress = float64(p) / float64(t) * 100
		}
		update := JobUpdate{
			Status:    "running",
			Progress:  progress,
			Processed: p,
			Total:     t,
		}
		job.Broadcaster.broadcast(update)
		_ = updateJobProgress(s.db, job.ID, "running", p, t, progress, "")
	}

	err := runFetch(job, report)

	var status string
	var errMsg string
	switch {
	case err == nil:
		status = "completed"
	case err.Error() == "cancelled":
		status = "cancelled"
	default:
		status = "error"
		errMsg = err.Error()
	}

	update := JobUpdate{
		Status:    status,
		Progress:  float64(totalToInt(total, processed)) / float64(max(total, 1)) * 100,
		Processed: processed,
		Total:     total,
		Error:     errMsg,
	}
	job.Broadcaster.broadcast(update)

	if errMsg != "" {
		log.Printf("Job %s failed: %s\n", job.ID, errMsg)
	}

	progress := 0.0
	if total > 0 {
		if status == "completed" {
			progress = 100
		} else {
			progress = float64(processed) / float64(total) * 100
		}
	}
	_ = updateJobProgress(s.db, job.ID, status, processed, total, progress, errMsg)
}

func (s *Server) StartScheduler() {
	go func() {
		ticker := time.NewTicker(60 * time.Second)
		defer ticker.Stop()
		for range ticker.C {
			s.tryScheduledRun()
		}
	}()
}

func (s *Server) tryScheduledRun() {
	enabled, _ := getSetting(s.db, "schedule_enabled")
	if enabled != "true" {
		return
	}

	lastRun, _ := getSetting(s.db, "last_scheduled_run")

	intervalHoursStr, _ := getSetting(s.db, "schedule_interval_hours")
	intervalDaysStr, _ := getSetting(s.db, "schedule_interval_days")

	intervalHours := parseIntOrZero(intervalHoursStr)
	intervalDays := parseIntOrZero(intervalDaysStr)
	totalInterval := time.Duration(intervalHours) * time.Hour
	if intervalDays > 0 {
		totalInterval = time.Duration(intervalDays) * 24 * time.Hour
	}

	if totalInterval <= 0 {
		return
	}

	if lastRun != "" {
		lastRunTime, err := time.Parse(time.RFC3339, lastRun)
		if err == nil && time.Since(lastRunTime) < totalInterval {
			return
		}
	}

	activeFileID, _ := getSetting(s.db, "active_file_id")
	if activeFileID == "" {
		return
	}

	upload, err := getUpload(s.db, activeFileID)
	if err != nil {
		return
	}

	paramsJSON, _ := getSetting(s.db, "last_fetch_params")
	if paramsJSON == "" {
		return
	}

	var req fetchRequest
	if err := json.Unmarshal([]byte(paramsJSON), &req); err != nil {
		return
	}

	req.UploadID = &activeFileID
	req.FileID = nil
	req.FilePath = nil

	jobID := uuid.New().String()
	ctx, cancel := context.WithCancel(context.Background())

	params := FetchParams{
		File:                  upload.FilePath,
		StartNumber:           req.StartNumber,
		EndNumber:             req.EndNumber,
		SheetName:             req.SheetName,
		SheetNames:            req.SheetNames,
		ArtikelnummerCol:      req.ArtikelnummerCol,
		PDFLinkCol:            req.PDFLinkCol,
		DocumentType:          req.DocumentType,
		UseDocumentTypeColumn: req.UseDocumentTypeColumn,
		DocumentTypeColumn:    req.DocumentTypeColumn,
		Company:               req.Company,
		UpdateDate:            req.UpdateDate,
		DateColumn:            req.DateColumn,
		Concurrency:           req.Concurrency,
	}

	paramsJSONBytes, _ := json.Marshal(req)
	_, _ = insertJob(s.db, jobID, upload.OriginalName, upload.FilePath, params, paramsJSONBytes)

	broadcaster := newJobBroadcaster()
	job := &PendingJob{
		ID:          jobID,
		Params:      params,
		FilePath:    upload.FilePath,
		FileName:    upload.OriginalName,
		Ctx:         ctx,
		Cancel:      cancel,
		Broadcaster: broadcaster,
	}

	s.mu.Lock()
	s.jobs[jobID] = job
	s.mu.Unlock()

	_ = setSetting(s.db, "last_scheduled_run", time.Now().UTC().Format(time.RFC3339))

	go s.runJob(job)
}

func parseIntOrZero(s string) int {
	var n int
	fmt.Sscanf(s, "%d", &n)
	return n
}

func totalToInt(total, processed int) int {
	if total > 0 {
		return total
	}
	return processed
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


