package main

// validateRowRange validates the row range against actual datapackage main

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"github.com/xuri/excelize/v2"
)

// Global variables with proper synchronization
var (
	cancelFunc context.CancelFunc
	cancelMu   sync.Mutex
)

// FetchDocumentsParams defines the parameters for document fetching
type FetchDocumentsParams struct {
	File         string `json:"file"`
	StartNumber  int    `json:"startNumber"`
	EndNumber    int    `json:"endNumber"`
	SheetName    string `json:"sheetName"`
	DocumentType string `json:"documentType"`
}

// DocumentTypes maps user-friendly names to API parameter values
var documentTypes = map[string]string{
	"Produktdatablad":   "TechnicalDataSheet",
	"Säkerhetsdatablad": "SafetySheets",
}

// initFetcher processes Excel file and fetches documents
func initFetcher(ctx context.Context, params FetchDocumentsParams) error {
	// Validate parameters
	if err := validateParams(params); err != nil {
		return fmt.Errorf("parameter validation failed: %w", err)
	}

	// Create cancellable context
	var cancelCtx context.Context
	cancelMu.Lock()
	cancelCtx, cancelFunc = context.WithCancel(ctx)
	cancelMu.Unlock()

	// Ensure cleanup
	defer func() {
		cancelMu.Lock()
		if cancelFunc != nil {
			cancelFunc()
			cancelFunc = nil
		}
		cancelMu.Unlock()
	}()

	// Open Excel file
	f, err := excelize.OpenFile(params.File)
	if err != nil {
		return fmt.Errorf("failed to open Excel file: %w", err)
	}
	defer f.Close()

	// Get rows from specified sheet
	rows, err := f.GetRows(params.SheetName)
	if err != nil {
		return fmt.Errorf("failed to read sheet '%s': %w", params.SheetName, err)
	}

	if len(rows) == 0 {
		return fmt.Errorf("sheet '%s' is empty", params.SheetName)
	}

	// Find required columns
	header := rows[0]
	artikelnummerCol, dokumenttypCol, err := findColumns(header)
	if err != nil {
		return err
	}

	// Find or create PDF link column
	pdfLinkCol := findOrCreatePDFColumn(f, params.SheetName, header)

	// Validate row range
	if err := validateRowRange(params, len(rows)); err != nil {
		return err
	}

	// Process rows
	if err := processRows(cancelCtx, f, rows, params, artikelnummerCol, dokumenttypCol, pdfLinkCol); err != nil {
		return err
	}

	// Save file if operation wasn't cancelled
	select {
	case <-cancelCtx.Done():
		runtime.EventsEmit(ctx, "progress_cancelled", "Operation cancelled by user")
		return fmt.Errorf("operation cancelled by user")
	default:
		if err := f.SaveAs(params.File); err != nil {
			return fmt.Errorf("failed to save Excel file: %w", err)
		}
	}

	runtime.EventsEmit(ctx, "progress_complete", "Document fetching completed successfully")
	return nil
}

// validateParams validates the input parameters
func validateParams(params FetchDocumentsParams) error {
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
	if _, exists := documentTypes[params.DocumentType]; !exists {
		return fmt.Errorf("invalid document type: %s", params.DocumentType)
	}
	return nil
}

// findColumns locates the required columns in the header
func findColumns(header []string) (int, int, error) {
	var artikelnummerCol, dokumenttypCol int = -1, -1

	for i, col := range header {
		col = strings.TrimSpace(col)
		switch col {
		case "Leverantörens artikelnummer":
			artikelnummerCol = i
		case "Dokumenttyp":
			dokumenttypCol = i
		}
	}

	if artikelnummerCol == -1 {
		return 0, 0, fmt.Errorf("required column 'Leverantörens artikelnummer' not found")
	}
	if dokumenttypCol == -1 {
		return 0, 0, fmt.Errorf("required column 'Dokumenttyp' not found")
	}

	return artikelnummerCol, dokumenttypCol, nil
}

// findOrCreatePDFColumn finds existing PDF column or creates a new one
func findOrCreatePDFColumn(f *excelize.File, sheetName string, header []string) int {
	// Look for the specific column "Filnamn eller webblänk"
	for i, col := range header {
		col = strings.TrimSpace(col)
		if col == "Filnamn eller webblänk" {
			return i
		}
	}

	// Look for other possible PDF/Link columns as fallback
	for i, col := range header {
		col = strings.TrimSpace(strings.ToLower(col))
		if strings.Contains(col, "pdf") || strings.Contains(col, "link") || strings.Contains(col, "dokumentlänk") || strings.Contains(col, "filnamn") || strings.Contains(col, "webblänk") {
			return i
		}
	}

	// Create new column for PDF links if none found
	pdfLinkCol := len(header)
	pdfLinkCell, _ := excelize.CoordinatesToCellName(pdfLinkCol+1, 1)
	f.SetCellValue(sheetName, pdfLinkCell, "Filnamn eller webblänk")

	return pdfLinkCol
}
func validateRowRange(params FetchDocumentsParams, totalRows int) error {
	if params.StartNumber > totalRows {
		return fmt.Errorf("start number (%d) exceeds total rows (%d)", params.StartNumber, totalRows)
	}
	if params.EndNumber > totalRows {
		return fmt.Errorf("end number (%d) exceeds total rows (%d)", params.EndNumber, totalRows)
	}
	return nil
}

// processRows processes the specified range of rows
func processRows(ctx context.Context, f *excelize.File, rows [][]string, params FetchDocumentsParams, artikelnummerCol, dokumenttypCol, pdfLinkCol int) error {
	startIndex := params.StartNumber - 1 // Convert to 0-based index
	endIndex := params.EndNumber

	// Ensure we don't exceed array bounds
	if endIndex > len(rows) {
		endIndex = len(rows)
	}

	totalRows := endIndex - startIndex
	docType := documentTypes[params.DocumentType]

	for i := startIndex; i < endIndex; i++ {
		select {
		case <-ctx.Done():
			runtime.EventsEmit(ctx, "progress_cancelled", "Operation cancelled during processing")
			return fmt.Errorf("operation cancelled")
		default:
			row := rows[i]

			// Skip if row doesn't have enough columns
			if len(row) <= artikelnummerCol {
				continue
			}

			artNum := strings.TrimSpace(row[artikelnummerCol])

			// Skip if article number is too short
			if len(artNum) < 5 {
				continue
			}

			// Add small delay to avoid overwhelming the server
			time.Sleep(100 * time.Millisecond)

			productCode := artNum[:5]
			pdfLink := fetchTechnicalDescription(productCode, docType)

			// Update progress
			progress := float64(i-startIndex+1) / float64(totalRows) * 100
			runtime.EventsEmit(ctx, "progress_update", progress)

			// Update Excel cells (1-based indexing)
			excelRowNum := i + 1

			// Update document type column
			dokumenttypCell, err := excelize.CoordinatesToCellName(dokumenttypCol+1, excelRowNum)
			if err != nil {
				return fmt.Errorf("failed to get cell coordinates: %w", err)
			}
			f.SetCellValue(params.SheetName, dokumenttypCell, params.DocumentType)

			// Update PDF link column
			pdfLinkCell, err := excelize.CoordinatesToCellName(pdfLinkCol+1, excelRowNum)
			if err != nil {
				return fmt.Errorf("failed to get cell coordinates: %w", err)
			}
			f.SetCellValue(params.SheetName, pdfLinkCell, pdfLink)
		}
	}

	return nil
}

// fetchTechnicalDescription fetches PDF link for a product code
func fetchTechnicalDescription(productCode, documentType string) string {

	// Multiple search strategies with decreasing specificity
	searchURLs := []string{
		fmt.Sprintf("https://www.illbruck.com/sv-se/teknisk-zon/teknisk-dokumentation/?search=%s&filters=type_%s", productCode, documentType),
		fmt.Sprintf("https://www.vandex.com/sv-se/teknisk-zon/teknisk-dokumentation/?search=%s&filters=type_%s", productCode[:4], documentType),
		fmt.Sprintf("https://www.nullifire.com/sv-se/teknisk-zon/teknisk-dokumentation/?search=%s&filters=type_%s", productCode[:3], documentType),
	}

	client := &http.Client{
		Timeout: 15 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: false,
		},
	}

	for i, url := range searchURLs {
		// Add delay between requests to be respectful
		if i > 0 {
			time.Sleep(200 * time.Millisecond)
		}

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			continue
		}

		// Set realistic headers
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
		req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
		req.Header.Set("Accept-Language", "sv-SE,sv;q=0.9,en;q=0.8")

		resp, err := client.Do(req)
		if err != nil {
			continue
		}

		// Process response and look for PDF links
		pdfLink := func() string {
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				return ""
			}

			doc, err := goquery.NewDocumentFromReader(resp.Body)
			if err != nil {
				return ""
			}

			// Look for PDF download links with multiple selectors
			selectors := []string{
				"a.downloads__action.align-vertical",
				"a[href*='.pdf']",
				".download-link[href*='.pdf']",
				"a.pdf-download",
			}

			for _, selector := range selectors {
				var foundLink string
				doc.Find(selector).Each(func(i int, s *goquery.Selection) {
					if foundLink != "" {
						return // Already found a link
					}
					if link, exists := s.Attr("href"); exists && strings.Contains(strings.ToLower(link), ".pdf") {
						// Make relative URLs absolute
						if strings.HasPrefix(link, "/") {
							link = "https://www.illbruck.com" + link
						}
						foundLink = link
					}
				})
				if foundLink != "" {
					return foundLink
				}
			}
			return ""
		}()

		// If we found a PDF link, return it
		if pdfLink != "" {
			return pdfLink
		}
	}

	return "Ingen PDF hittad..."
}

// cancelFetcher cancels the current fetch operation
func cancelFetcher() {
	cancelMu.Lock()
	defer cancelMu.Unlock()

	if cancelFunc != nil {
		cancelFunc()
		cancelFunc = nil
	}
}
