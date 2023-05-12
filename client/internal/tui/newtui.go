package tui

import (
	"errors"
	"time"

	"github.com/RomanIkonnikov93/keeper/client/internal/gapi"
	"github.com/RomanIkonnikov93/keeper/client/internal/models"

	"github.com/rivo/tview"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TUI struct {
	*tview.Application
	pages  *tview.Pages
	client *gapi.KeeperServiceClient
}

func NewTUI(client *gapi.KeeperServiceClient) *TUI {

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

func (t *TUI) StartScanningChanges() {

	// Goroutine for first scan.
	go func() {
		for {
			if t.client.Auth.Token != "" {
				err := t.client.CheckChanges()
				if errors.Is(err, status.Error(codes.Unauthenticated, "")) {
					t.authPage("Unauthenticated")
					return
				}
				if err != nil {
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
						t.authPage("Unauthenticated")
						return
					}
					if err != nil {
						t.authPage("Unknown error")
						return
					}
				}
			}
		}
	}()
}
