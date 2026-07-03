package main

import (
	"database/sql"
	"encoding/json"
	"os"
	"path/filepath"
	"time"

	_ "modernc.org/sqlite"
)

func initDB(path string) (*sql.DB, error) {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(1)

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS uploads (
		id TEXT PRIMARY KEY,
		file_name TEXT NOT NULL,
		original_name TEXT NOT NULL,
		file_path TEXT NOT NULL,
		uploaded_at TEXT NOT NULL DEFAULT (datetime('now'))
	)`)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS settings (
		key TEXT PRIMARY KEY,
		value TEXT NOT NULL DEFAULT ''
	)`)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS jobs (
		id TEXT PRIMARY KEY,
		file_name TEXT NOT NULL,
		file_path TEXT NOT NULL,
		sheet_name TEXT NOT NULL,
		company TEXT NOT NULL,
		document_type TEXT NOT NULL DEFAULT '',
		start_number INTEGER NOT NULL,
		end_number INTEGER NOT NULL,
		artikelnummer_col TEXT NOT NULL,
		pdf_link_col TEXT NOT NULL,
		use_document_type_column INTEGER NOT NULL DEFAULT 0,
		document_type_column TEXT NOT NULL DEFAULT '',
		update_date INTEGER NOT NULL DEFAULT 0,
		date_column TEXT NOT NULL DEFAULT '',
		status TEXT NOT NULL DEFAULT 'pending',
		total_rows INTEGER NOT NULL DEFAULT 0,
		processed_rows INTEGER NOT NULL DEFAULT 0,
		progress REAL NOT NULL DEFAULT 0.0,
		error_message TEXT NOT NULL DEFAULT '',
		params_json TEXT NOT NULL DEFAULT '',
		created_at TEXT NOT NULL DEFAULT (datetime('now')),
		updated_at TEXT NOT NULL DEFAULT (datetime('now'))
	)`)
	if err != nil {
		return nil, err
	}

	return db, nil
}

type jobParamsJSON struct {
	FileID                 *string `json:"file_id,omitempty"`
	FilePath               *string `json:"file_path,omitempty"`
	Name                   string  `json:"name"`
	SheetName              string  `json:"sheet_name"`
	Company                string  `json:"company"`
	DocumentType           string  `json:"document_type"`
	StartNumber            int     `json:"start_number"`
	EndNumber              int     `json:"end_number"`
	ArtikelnummerCol       string  `json:"artikelnummer_col"`
	PDFLinkCol             string  `json:"pdf_link_col"`
	UseDocumentTypeColumn  bool    `json:"use_document_type_column"`
	DocumentTypeColumn     string  `json:"document_type_column"`
	UpdateDate             bool    `json:"update_date"`
	DateColumn             string  `json:"date_column"`
}

type jobRow struct {
	ID                  string  `json:"id"`
	FileName            string  `json:"file_name"`
	FilePath            string  `json:"file_path"`
	SheetName           string  `json:"sheet_name"`
	Company             string  `json:"company"`
	Status              string  `json:"status"`
	TotalRows           int     `json:"total_rows"`
	ProcessedRows       int     `json:"processed_rows"`
	Progress            float64 `json:"progress"`
	ErrorMessage        string  `json:"error_message"`
	CreatedAt           string  `json:"created_at"`
	UpdatedAt           string  `json:"updated_at"`
}

func insertJob(db *sql.DB, id, fileName, filePath string, p FetchParams, paramsJSON []byte) (*jobRow, error) {
	row := &jobRow{
		ID:       id,
		FileName: fileName,
		FilePath: filePath,
	}
	_, err := db.Exec(`INSERT INTO jobs (
		id, file_name, file_path, sheet_name, company,
		document_type, start_number, end_number,
		artikelnummer_col, pdf_link_col,
		use_document_type_column, document_type_column,
		update_date, date_column,
		params_json, status
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, 'running')`,
		id, fileName, filePath,
		p.SheetName, p.Company,
		p.DocumentType, p.StartNumber, p.EndNumber,
		p.ArtikelnummerCol, p.PDFLinkCol,
		btoi(p.UseDocumentTypeColumn), p.DocumentTypeColumn,
		btoi(p.UpdateDate), p.DateColumn,
		string(paramsJSON),
	)
	if err != nil {
		return nil, err
	}
	return row, nil
}

func updateJobProgress(db *sql.DB, id, status string, processed, total int, progress float64, errMsg string) error {
	_, err := db.Exec(`UPDATE jobs SET
		status = ?, processed_rows = ?, total_rows = ?, progress = ?, error_message = ?,
		updated_at = datetime('now')
		WHERE id = ?`,
		status, processed, total, progress, errMsg, id,
	)
	return err
}

func listJobs(db *sql.DB) ([]jobRow, error) {
	rows, err := db.Query(`SELECT
		id, file_name, file_path, sheet_name, company,
		status, total_rows, processed_rows, progress, error_message,
		created_at, updated_at
		FROM jobs ORDER BY created_at DESC LIMIT 100`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jobs []jobRow
	for rows.Next() {
		var j jobRow
		if err := rows.Scan(&j.ID, &j.FileName, &j.FilePath,
			&j.SheetName, &j.Company,
			&j.Status, &j.TotalRows, &j.ProcessedRows, &j.Progress, &j.ErrorMessage,
			&j.CreatedAt, &j.UpdatedAt,
		); err != nil {
			return nil, err
		}
		jobs = append(jobs, j)
	}
	return jobs, rows.Err()
}

func getJobRow(db *sql.DB, id string) (*jobRow, error) {
	var j jobRow
	err := db.QueryRow(`SELECT
		id, file_name, file_path, sheet_name, company,
		status, total_rows, processed_rows, progress, error_message,
		created_at, updated_at
		FROM jobs WHERE id = ?`, id).Scan(
		&j.ID, &j.FileName, &j.FilePath,
		&j.SheetName, &j.Company,
		&j.Status, &j.TotalRows, &j.ProcessedRows, &j.Progress, &j.ErrorMessage,
		&j.CreatedAt, &j.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &j, nil
}

func getJobParamsJSON(db *sql.DB, id string) ([]byte, error) {
	var raw string
	err := db.QueryRow(`SELECT params_json FROM jobs WHERE id = ?`, id).Scan(&raw)
	if err != nil {
		return nil, err
	}
	return []byte(raw), nil
}

func mustMarshal(v interface{}) []byte {
	b, _ := json.Marshal(v)
	return b
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func jsonTimeNow() string {
	return time.Now().UTC().Format(time.RFC3339)
}

type uploadRow struct {
	ID           string `json:"id"`
	FileName     string `json:"file_name"`
	OriginalName string `json:"original_name"`
	FilePath     string `json:"file_path"`
	UploadedAt   string `json:"uploaded_at"`
}

func insertUpload(db *sql.DB, id, originalName, filePath string) error {
	_, err := db.Exec(`INSERT INTO uploads (id, file_name, original_name, file_path) VALUES (?, ?, ?, ?)`,
		id, originalName, originalName, filePath)
	return err
}

func listUploads(db *sql.DB) ([]uploadRow, error) {
	rows, err := db.Query(`SELECT id, file_name, original_name, file_path, uploaded_at FROM uploads ORDER BY uploaded_at DESC LIMIT 100`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var uploads []uploadRow
	for rows.Next() {
		var u uploadRow
		if err := rows.Scan(&u.ID, &u.FileName, &u.OriginalName, &u.FilePath, &u.UploadedAt); err != nil {
			return nil, err
		}
		uploads = append(uploads, u)
	}
	if uploads == nil {
		uploads = []uploadRow{}
	}
	return uploads, rows.Err()
}

func getUpload(db *sql.DB, id string) (*uploadRow, error) {
	var u uploadRow
	err := db.QueryRow(`SELECT id, file_name, original_name, file_path, uploaded_at FROM uploads WHERE id = ?`, id).Scan(
		&u.ID, &u.FileName, &u.OriginalName, &u.FilePath, &u.UploadedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func deleteUpload(db *sql.DB, id string) error {
	_, err := db.Exec(`DELETE FROM uploads WHERE id = ?`, id)
	return err
}

func getSetting(db *sql.DB, key string) (string, error) {
	var val string
	err := db.QueryRow(`SELECT value FROM settings WHERE key = ?`, key).Scan(&val)
	if err == sql.ErrNoRows {
		return "", nil
	}
	return val, err
}

func setSetting(db *sql.DB, key, value string) error {
	_, err := db.Exec(`INSERT INTO settings (key, value) VALUES (?, ?) ON CONFLICT(key) DO UPDATE SET value = excluded.value`, key, value)
	return err
}

func getAllSettings(db *sql.DB) (map[string]string, error) {
	rows, err := db.Query(`SELECT key, value FROM settings`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	settings := make(map[string]string)
	for rows.Next() {
		var k, v string
		if err := rows.Scan(&k, &v); err != nil {
			return nil, err
		}
		settings[k] = v
	}
	return settings, rows.Err()
}
