package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/xuri/excelize/v2"
)

type FetchParams struct {
	File                  string   `json:"file"`
	StartNumber           int      `json:"startNumber"`
	EndNumber             int      `json:"endNumber"`
	SheetName             string   `json:"sheetName"`
	SheetNames            []string `json:"sheetNames,omitempty"`
	ArtikelnummerCol      string   `json:"artikelnummerCol"`
	PDFLinkCol            string   `json:"pdfLinkCol"`
	DocumentType          string   `json:"documentType"`
	UseDocumentTypeColumn bool     `json:"use_document_type_column"`
	DocumentTypeColumn    string   `json:"document_type_column"`
	Company               string   `json:"company"`
	UpdateDate            bool     `json:"update_date"`
	DateColumn            string   `json:"date_column"`
	Concurrency           int      `json:"concurrency"`
}

var documentTypes = map[string]string{
	"Produktdatablad":            "TechnicalDataSheet",
	"Säkerhetsdatablad":          "SafetySheets",
	"Prestandadeklaration":       "Declaration",
	"Miljövarudeklaration (EPD)": "EnvironmentalProductDeclaration",
	"Certifikat":                 "TestCertificationDocument",
}

var companyURLs = map[string]string{
	"illbruck":  "https://www.illbruck.com/sv-se/teknisk-zon/teknisk-dokumentation/?search=%s&filters=type_%s",
	"vandex":    "https://www.vandex.com/sv-se/teknisk-zon/teknisk-dokumentation/?search=%s&filters=type_%s",
	"nullifire": "https://www.nullifire.com/sv-se/teknisk-zon/teknisk-dokumentation/?search=%s&filters=type_%s",
	"flowcrete": "https://www.flowcrete.eu/sv-se/produkter-golvsystem/soek-produkter-och-golvsystem/?search=%s&filters=type_%s",
	"matacryl":  "https://www.tremco-europe.com/sv-se/produkter-system/soek-produkter/?search=%s&filters=type_%s",
	"tremco":    "https://www.tremcocpg.eu/sv-se/produkter-loesningar/teknisk-dokumentation/?search=%s&filters=type_%s",
}

type progressCB func(processed, total int)

type rowTask struct {
	excelRow int
	sheet    string
	artNum   string
	docType  string
}

type sheetWrite struct {
	pdfLinkIdx int
	dateIdx    int
	tasks      []rowTask
}

func resolveSheets(params FetchParams, allSheets []string) ([]string, error) {
	if len(params.SheetNames) > 0 {
		for _, name := range params.SheetNames {
			found := false
			for _, s := range allSheets {
				if s == name {
					found = true
					break
				}
			}
			if !found {
				return nil, fmt.Errorf("sheet '%s' not found", name)
			}
		}
		return params.SheetNames, nil
	}
	if params.SheetName != "" {
		for _, s := range allSheets {
			if s == params.SheetName {
				return []string{params.SheetName}, nil
			}
		}
		return nil, fmt.Errorf("sheet '%s' not found", params.SheetName)
	}
	return nil, fmt.Errorf("no sheet specified")
}

func runFetch(job *PendingJob, report progressCB) error {
	params := job.Params

	if err := validateParams(params); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	concurrency := params.Concurrency
	if concurrency <= 0 {
		concurrency = 5
	}

	f, err := excelize.OpenFile(job.FilePath)
	if err != nil {
		return fmt.Errorf("failed to open Excel file: %w", err)
	}
	defer f.Close()

	allSheets := f.GetSheetList()

	sheets, err := resolveSheets(params, allSheets)
	if err != nil {
		return err
	}

	searchURL, ok := companyURLs[params.Company]
	if !ok {
		return fmt.Errorf("unknown company: %s", params.Company)
	}

	startIndex := params.StartNumber - 1

	type sheetData struct {
		write  sheetWrite
		header []string
	}

	var sheetDataMap []sheetData
	var allTasks []rowTask
	grandTotal := 0

	for _, sheetName := range sheets {
		rows, err := f.GetRows(sheetName)
		if err != nil {
			return fmt.Errorf("failed to read sheet '%s': %w", sheetName, err)
		}
		if len(rows) == 0 {
			return fmt.Errorf("sheet '%s' is empty", sheetName)
		}

		header := rows[0]
		artikelnummerIdx, pdfLinkIdx, err := findColumnsDynamic(header, params)
		if err != nil {
			return err
		}

		if pdfLinkIdx == -1 {
			pdfLinkIdx = createPDFColumn(f, sheetName, header)
		}

		var docTypeColIdx int = -1
		if params.UseDocumentTypeColumn {
			for i, col := range header {
				if strings.TrimSpace(col) == params.DocumentTypeColumn {
					docTypeColIdx = i
					break
				}
			}
			if docTypeColIdx == -1 {
				return fmt.Errorf("document type column '%s' not found in sheet '%s'", params.DocumentTypeColumn, sheetName)
			}
		}

		manualDocType := ""
		if !params.UseDocumentTypeColumn {
			manualDocType = documentTypes[params.DocumentType]
		}

		dateIdx := -1
		if params.UpdateDate && params.DateColumn != "" {
			dateIdx, _ = findDateColumn(header, params.DateColumn)
		}

		endIndex := params.EndNumber
		if endIndex > len(rows) {
			endIndex = len(rows)
		}

		var sheetTasks []rowTask
		for i := startIndex; i < endIndex; i++ {
			select {
			case <-job.Ctx.Done():
				return fmt.Errorf("cancelled")
			default:
			}

			row := rows[i]
			if len(row) <= artikelnummerIdx {
				continue
			}
			artNum := strings.TrimSpace(row[artikelnummerIdx])
			if len(artNum) < 5 {
				continue
			}

			var docType string
			if params.UseDocumentTypeColumn {
				if len(row) <= docTypeColIdx {
					continue
				}
				excelDocType := strings.TrimSpace(row[docTypeColIdx])
				mapped, ok := documentTypes[excelDocType]
				if !ok {
					log.Printf("Row %d in '%s' has unknown document type '%s'\n", i+1, sheetName, excelDocType)
					continue
				}
				docType = mapped
			} else {
				docType = manualDocType
			}

			sheetTasks = append(sheetTasks, rowTask{
				excelRow: i + 1,
				sheet:    sheetName,
				artNum:   artNum,
				docType:  docType,
			})
		}

		sheetDataMap = append(sheetDataMap, sheetData{
			write: sheetWrite{
				pdfLinkIdx: pdfLinkIdx,
				dateIdx:    dateIdx,
				tasks:      sheetTasks,
			},
			header: header,
		})

		allTasks = append(allTasks, sheetTasks...)
		grandTotal += len(sheetTasks)
	}

	if grandTotal == 0 {
		return fmt.Errorf("no rows to process")
	}

	todayStr := time.Now().Format("2006-01-02")

	if concurrency > grandTotal {
		concurrency = grandTotal
	}

	taskCh := make(chan rowTask, grandTotal)
	resultCh := make(chan rowTask, grandTotal)

	var wg sync.WaitGroup
	for w := 0; w < concurrency; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range taskCh {
				select {
				case <-job.Ctx.Done():
					return
				default:
				}
				productCode := task.artNum[:5]
				pdfLink := fetchTechnicalDescription(searchURL, productCode, task.docType)
				select {
				case resultCh <- rowTask{
					excelRow: task.excelRow,
					sheet:    task.sheet,
					artNum:   pdfLink,
					docType:  "",
				}:
				case <-job.Ctx.Done():
					return
				}
			}
		}()
	}

	go func() {
		for _, t := range allTasks {
			select {
			case taskCh <- t:
			case <-job.Ctx.Done():
				break
			}
		}
		close(taskCh)
	}()

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	sheetTaskIndex := make(map[string]int)
	for sheetIdx, sd := range sheetDataMap {
		for _, t := range sd.write.tasks {
			key := fmt.Sprintf("%s_%d", t.sheet, t.excelRow)
			sheetTaskIndex[key] = sheetIdx
		}
	}

	processedCount := 0
loop:
	for {
		select {
		case <-job.Ctx.Done():
			return fmt.Errorf("cancelled")
		case result, ok := <-resultCh:
			if !ok {
				break loop
			}
			key := fmt.Sprintf("%s_%d", result.sheet, result.excelRow)
			sheetIdx := sheetTaskIndex[key]
			sd := &sheetDataMap[sheetIdx]

			pdfCell, _ := excelize.CoordinatesToCellName(sd.write.pdfLinkIdx+1, result.excelRow)
			f.SetCellValue(result.sheet, pdfCell, result.artNum)

			if sd.write.dateIdx >= 0 {
				dateCell, _ := excelize.CoordinatesToCellName(sd.write.dateIdx+1, result.excelRow)
				f.SetCellValue(result.sheet, dateCell, todayStr)
			}

			processedCount++
			report(processedCount, grandTotal)
		}
	}

	if err := f.SaveAs(job.FilePath); err != nil {
		return fmt.Errorf("failed to save file: %w", err)
	}

	return nil
}

func findColumnsDynamic(header []string, params FetchParams) (int, int, error) {
	var artikelnummerIdx, pdfLinkIdx int = -1, -1
	for i, col := range header {
		colTrim := strings.TrimSpace(col)
		if colTrim == params.ArtikelnummerCol {
			artikelnummerIdx = i
		}
		if params.PDFLinkCol != "" && colTrim == params.PDFLinkCol {
			pdfLinkIdx = i
		}
	}
	if artikelnummerIdx == -1 {
		return 0, 0, fmt.Errorf("article number column '%s' not found", params.ArtikelnummerCol)
	}
	return artikelnummerIdx, pdfLinkIdx, nil
}

func createPDFColumn(f *excelize.File, sheetName string, header []string) int {
	colIdx := len(header)
	cell, _ := excelize.CoordinatesToCellName(colIdx+1, 1)
	f.SetCellValue(sheetName, cell, "Filnamn eller webblänk")
	return colIdx
}

func findDateColumn(header []string, dateColumn string) (int, error) {
	for i, col := range header {
		if strings.TrimSpace(col) == dateColumn {
			return i, nil
		}
	}
	return -1, fmt.Errorf("date column '%s' not found", dateColumn)
}

func validateParams(params FetchParams) error {
	if params.File == "" {
		return fmt.Errorf("file path is required")
	}
	if params.SheetName == "" && len(params.SheetNames) == 0 {
		return fmt.Errorf("sheet name is required")
	}
	if params.StartNumber < 1 {
		return fmt.Errorf("start number must be >= 1")
	}
	if params.EndNumber < params.StartNumber {
		return fmt.Errorf("end number must be >= start number")
	}
	if !params.UseDocumentTypeColumn {
		if params.DocumentType == "" {
			return fmt.Errorf("document type required unless using type column")
		}
		if _, ok := documentTypes[params.DocumentType]; !ok {
			return fmt.Errorf("invalid document type")
		}
	}
	return nil
}

func fetchTechnicalDescription(url, productCode, documentType string) string {
	searchURL := fmt.Sprintf(url, productCode, documentType)

	client := &http.Client{Timeout: 15 * time.Second}
	req, _ := http.NewRequest("GET", searchURL, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0")
	resp, err := client.Do(req)
	if err != nil {
		return "Ingen PDF hittad..."
	}

	pdf := extractPDFLink(resp)
	if pdf != "" {
		return pdf
	}
	return "Ingen PDF hittad..."
}

func extractPDFLink(resp *http.Response) string {
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return ""
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return ""
	}
	selectors := []string{"a.downloads__action.align-vertical", "a[href*='.pdf']", ".download-link[href*='.pdf']", "a.pdf-download"}
	for _, sel := range selectors {
		var found string
		doc.Find(sel).Each(func(i int, s *goquery.Selection) {
			if found != "" {
				return
			}
			if link, ok := s.Attr("href"); ok && strings.Contains(strings.ToLower(link), ".pdf") {
				if strings.HasPrefix(link, "/") {
					link = "https://www.illbruck.com" + link
				}
				found = link
			}
		})
		if found != "" {
			return found
		}
	}
	return ""
}
