package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"github.com/xuri/excelize/v2"
)

var (
	cancelFunc context.CancelFunc
	cancelMu   sync.Mutex
)

type FetchParams struct {
	File                  string `json:"file"`
	StartNumber           int    `json:"startNumber"`
	EndNumber             int    `json:"endNumber"`
	SheetName             string `json:"sheetName"`
	ArtikelnummerCol      string `json:"artikelnummerCol"`
	PDFLinkCol            string `json:"pdfLinkCol"`
	DocumentType          string `json:"documentType"`
	UseDocumentTypeColumn bool   `json:"use_document_type_column"`
	DocumentTypeColumn    string `json:"document_type_column"`
	Company               string `json:"company"`
	UpdateDate            bool   `json:"update_date"`
	DateColumn            string `json:"date_column"`
}

var documentTypes = map[string]string{
	"Produktdatablad":            "TechnicalDataSheet",
	"Säkerhetsdatablad":          "SafetySheets",
	"Prestandadeklaration":       "Declaration",
	"Miljövarudeklaration (EPD)": "EnvironmentalProductDeclaration",
	"Certifikat":                 "TestCertificationDocument",
}

var company = map[string]string{
	"illbruck":  "https://www.illbruck.com/sv-se/teknisk-zon/teknisk-dokumentation/?search=%s&filters=type_%s",
	"vandex":    "https://www.vandex.com/sv-se/teknisk-zon/teknisk-dokumentation/?search=%s&filters=type_%s",
	"nullifire": "https://www.nullifire.com/sv-se/teknisk-zon/teknisk-dokumentation/?search=%s&filters=type_%s",
	"flowcrete": "https://www.flowcrete.eu/sv-se/produkter-golvsystem/soek-produkter-och-golvsystem/?search=%s&filters=type_%s",
	"matacryl":  "https://www.tremco-europe.com/sv-se/produkter-system/soek-produkter/?search=%s&filters=type_%s",
	"tremco":    "https://www.tremcocpg.eu/sv-se/produkter-loesningar/teknisk-dokumentation/?search=%s&filters=type_%s",
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

func fetchTechnicalDescription(url, productCode, documentType string) string {

	searchUrl := fmt.Sprintf(url, productCode, documentType)

	client := &http.Client{Timeout: 15 * time.Second}

	req, _ := http.NewRequest("GET", searchUrl, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0")
	resp, _ := client.Do(req)

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

func setupCancellableContext(ctx context.Context) context.Context {
	cancelMu.Lock()
	defer cancelMu.Unlock()
	cancelCtx, cancel := context.WithCancel(ctx)
	cancelFunc = cancel
	return cancelCtx
}

func cleanupContext() {
	cancelMu.Lock()
	defer cancelMu.Unlock()
	if cancelFunc != nil {
		cancelFunc()
		cancelFunc = nil
	}
}

func validateParams(params FetchParams) error {
	if params.File == "" {
		return fmt.Errorf("file path is required")
	}
	if params.SheetName == "" {
		return fmt.Errorf("sheet name is required")
	}
	if params.StartNumber < 1 {
		return fmt.Errorf("start number must be >= 1")
	}
	if params.EndNumber < params.StartNumber {
		return fmt.Errorf("end number must be >= start number")
	}
	// Only validate manual mode
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

func (a *App) processFetch(params FetchParams) error {
	if err := validateParams(params); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	cancelCtx := setupCancellableContext(a.ctx)
	defer cleanupContext()

	f, err := excelize.OpenFile(params.File)
	if err != nil {
		return fmt.Errorf("failed to open Excel file: %w", err)
	}
	defer f.Close()

	sheetNames := f.GetSheetList()
	foundSheet := false
	for _, s := range sheetNames {
		if s == params.SheetName {
			foundSheet = true
			break
		}
	}
	if !foundSheet {
		return fmt.Errorf("sheet '%s' not found", params.SheetName)
	}

	rows, err := f.GetRows(params.SheetName)
	if err != nil {
		return fmt.Errorf("failed to read sheet '%s': %w", params.SheetName, err)
	}
	if len(rows) == 0 {
		return fmt.Errorf("sheet '%s' is empty", params.SheetName)
	}

	header := rows[0]
	artikelnummerIdx, pdfLinkIdx, err := findColumnsDynamic(header, params)
	if err != nil {
		return err
	}

	if pdfLinkIdx == -1 {
		pdfLinkIdx = createPDFColumn(f, params.SheetName, header)
	}

	// Find document type column if needed
	var docTypeColIdx int = -1
	if params.UseDocumentTypeColumn {
		for i, col := range header {
			if strings.TrimSpace(col) == params.DocumentTypeColumn {
				docTypeColIdx = i
				break
			}
		}
		if docTypeColIdx == -1 {
			return fmt.Errorf("document type column '%s' not found", params.DocumentTypeColumn)
		}
	}

	// Manual type mapping
	manualDocType := ""
	if !params.UseDocumentTypeColumn {
		manualDocType = documentTypes[params.DocumentType]
	}

	startIndex := params.StartNumber - 1
	endIndex := params.EndNumber
	if endIndex > len(rows) {
		endIndex = len(rows)
	}

	totalRows := endIndex - startIndex

	searchCompany, ok := company[params.Company]
	if !ok {
		return fmt.Errorf("no company found")
	}

	processedCount := 0
	for i := startIndex; i < endIndex; i++ {
		select {
		case <-cancelCtx.Done():
			runtime.EventsEmit(a.ctx, "progress_cancelled", "Operation cancelled")
			return fmt.Errorf("operation cancelled")
		default:
			row := rows[i]

			if len(row) <= artikelnummerIdx {
				continue
			}

			artNum := strings.TrimSpace(row[artikelnummerIdx])
			if len(artNum) < 5 {
				continue
			}

			// Determine document type
			var docType string
			if params.UseDocumentTypeColumn {
				if len(row) <= docTypeColIdx {
					continue
				}
				excelDocType := strings.TrimSpace(row[docTypeColIdx])
				mapped, ok := documentTypes[excelDocType]
				if !ok {
					log.Printf("Row %d has unknown document type '%s'\n", i+1, excelDocType)
					continue
				}
				docType = mapped
			} else {
				docType = manualDocType
			}

			productCode := artNum[:5]
			pdfLink := fetchTechnicalDescription(searchCompany, productCode, docType)

			excelRow := i + 1
			pdfCell, _ := excelize.CoordinatesToCellName(pdfLinkIdx+1, excelRow)
			f.SetCellValue(params.SheetName, pdfCell, pdfLink)

			// Update date column
			if params.UpdateDate && params.DateColumn != "" {
				dateIdx, err := findDateColumn(header, params.DateColumn)
				if err == nil {
					dateCell, _ := excelize.CoordinatesToCellName(dateIdx+1, excelRow)
					f.SetCellValue(params.SheetName, dateCell, time.Now().Format("2006-01-02"))
				}
			}

			processedCount++
			progress := float64(i-startIndex+1) / float64(totalRows) * 100
			runtime.EventsEmit(a.ctx, "progress_update", progress)
			time.Sleep(100 * time.Millisecond)
		}
	}

	if err := f.SaveAs(params.File); err != nil {
		return fmt.Errorf("failed to save file: %w", err)
	}

	runtime.EventsEmit(a.ctx, "progress_complete",
		fmt.Sprintf("Completed! Processed %d of %d rows", processedCount, totalRows))

	return nil
}

func findDateColumn(header []string, dateColumn string) (int, error) {
	for i, col := range header {
		if strings.TrimSpace(col) == dateColumn {
			return i, nil
		}
	}
	return -1, fmt.Errorf("date column '%s' not found", dateColumn)
}
