package main

import (
	"context"
	"fmt"
	"log"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// SelectExcelFile opens a file dialog for selecting Excel files
func (a *App) SelectExcelFile() (string, error) {
	filePath, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Välj Excel-fil",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Excel Files (*.xlsx)",
				Pattern:     "*.xlsx",
			},
			{
				DisplayName: "All Files (*.*)",
				Pattern:     "*.*",
			},
		},
	})

	if err != nil {
		log.Printf("Error opening file dialog: %v\n", err)
		return "", err
	}

	log.Printf("Selected file: %s\n", filePath)
	return filePath, nil
}

// FetchDocuments is the main entry point called from the frontend
func (a *App) FetchDocuments(params FetchParams) error {
	log.Printf("FetchDocuments called with params: %+v\n", params)

	// Emit start event
	runtime.EventsEmit(a.ctx, "progress_start", "Starting document fetch...")

	err := a.processFetch(params)
	if err != nil {
		log.Printf("Error in FetchDocuments: %v\n", err)
		runtime.EventsEmit(a.ctx, "progress_error", err.Error())
		return err
	}

	log.Println("FetchDocuments completed successfully")
	return nil
}

//func (a *App) FetchImages(params FetchImagesParams) {
//	initImageFetcher(a.ctx, params)
//}

func (a *App) CancelFetch() {
	log.Println("Cancel requested from frontend")
	cancelMu.Lock()
	if cancelFunc != nil {
		cancelFunc()
	}
	cancelMu.Unlock()
}
