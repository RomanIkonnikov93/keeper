package tui

import (
	"errors"
	"time"

	"github.com/RomanIkonnikov93/keeper/client/internal/grpcapi"
	"github.com/RomanIkonnikov93/keeper/client/internal/models"

	"github.com/rivo/tview"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TUI is a struct for terminal user interface.
type TUI struct {
	*tview.Application
	pages  *tview.Pages
	client *grpcapi.KeeperServiceClient
}

// NewTUI gets new terminal user interface for client.
func NewTUI(client *grpcapi.KeeperServiceClient) *TUI {

	app := tview.NewApplication()
	pages := tview.NewPages()

	app.SetRoot(pages, true)

	t := &TUI{
		Application: app,
		pages:       pages,
		client:      client,
	}

	t.StartScanningChanges()

	t.authPage("")

	return t
}

// StartScanningChanges starts a loop to check for new records on the server.
func (t *TUI) StartScanningChanges() {

	// Goroutine for first scan.
	go func() {
		for {
			if t.client.Auth.Token != "" {
				err := t.client.CheckChanges()
				if errors.Is(err, status.Error(codes.Unauthenticated, "")) {
					t.client.Logger.Error(err)
					t.authPage("Unauthenticated")
					return
				}
				if err != nil {
					t.client.Logger.Error(err)
					t.authPage("Unknown error")
					return
				}
				break
			}
		}
	}()

	go func() {
		ticker := time.NewTicker(models.Ticker)
		for {
			select {
			case <-ticker.C:
				if t.client.Auth.Token != "" {
					err := t.client.CheckChanges()
					if errors.Is(err, status.Error(codes.Unauthenticated, "")) {
						t.client.Logger.Error(err)
						t.authPage("Unauthenticated")
						return
					}
					if err != nil {
						t.client.Logger.Error(err)
						t.authPage("Unknown error")
						return
					}
				}
			}
		}
	}()
}
