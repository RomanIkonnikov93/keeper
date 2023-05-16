package tui

import (
	"errors"
	"strconv"

	"github.com/RomanIkonnikov93/keeper/client/internal/models"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// filePage switches to file page.
func (t *TUI) filePage(message string) {

	switch t.client.Record.ActionType {
	case models.Add:
		form := tview.NewForm().
			AddInputField("Decryption:", "", 30, nil, func(description string) {
				t.client.Record.Description = description
			}).
			AddTextView("", "Max file size: 2Mb", 100, 1, true, false).
			AddInputField("File path:", "", 60, nil, func(path string) {
				t.client.Record.FilePath = path
			}).
			AddButton("Add", func() {

				err := t.client.AddRecord()

				if errors.Is(err, status.Error(codes.InvalidArgument, "")) {
					t.mainPage("InvalidArgument")
					return
				}
				if errors.Is(err, models.ErrMaxFileSize) {
					t.mainPage("file size exceeded")
					return
				}
				if err != nil {
					t.mainPage("Unknown error")
					return
				}

				t.mainPage("OK")
			}).
			AddButton("Back", func() {
				t.mainPage("")
			})

		frame := tview.NewFrame(form).SetBorders(0, 0, 0, 1, 4, 4).
			AddText("TAB - for switching between fields | Enter - for select", false, tview.AlignLeft, tcell.ColorWhite).
			AddText(message, false, tview.AlignRight, tcell.ColorRed)

		t.pages.AddPage("file", frame, true, true)
		t.pages.SwitchToPage("file")

	case models.Get:
		form := tview.NewForm().
			AddInputField("Record ID:", "", 30, nil, func(id string) {
				res, err := strconv.Atoi(id)
				if err != nil {
					t.filePage("invalid id")
					return
				}
				t.client.Record.RecordID = int32(res)
			}).
			AddTextView("Default path:", t.client.Cfg.DownloadFilesPath, 100, 1, true, false).
			AddInputField("Set download path:", "", 60, nil, func(path string) {
				t.client.Record.FilePath = path
			}).
			AddButton("Get", func() {

				err := t.client.GetRecordByID()

				if errors.Is(err, status.Error(codes.InvalidArgument, "")) {
					t.mainPage("InvalidArgument")
					return
				}
				if errors.Is(err, models.ErrNotExist) {
					t.mainPage("record not exist")
					return
				}
				if err != nil {
					t.mainPage("Unknown error")
					return
				}

				t.outputPage("OK")
			}).
			AddButton("Back", func() {
				t.mainPage("")
			})

		frame := tview.NewFrame(form).SetBorders(0, 0, 0, 1, 4, 4).
			AddText("TAB - for switching between fields | Enter - for select", false, tview.AlignLeft, tcell.ColorWhite).
			AddText(message, false, tview.AlignRight, tcell.ColorRed)

		t.pages.AddPage("file", frame, true, true)
		t.pages.SwitchToPage("file")

	case models.Update:
		form := tview.NewForm().
			AddInputField("Record ID:", "", 30, nil, func(id string) {
				res, err := strconv.Atoi(id)
				if err != nil {
					t.filePage("invalid id")
					return
				}
				t.client.Record.RecordID = int32(res)
			}).
			AddInputField("Decryption:", "", 30, nil, func(description string) {
				t.client.Record.Description = description
			}).
			AddTextView("", "Max file size: 2Mb", 100, 1, true, false).
			AddInputField("File path:", "", 60, nil, func(path string) {
				t.client.Record.FilePath = path
			}).
			AddButton("Update", func() {

				err := t.client.UpdateRecordByID()

				if errors.Is(err, status.Error(codes.InvalidArgument, "")) {
					t.mainPage("InvalidArgument")
					return
				}
				if errors.Is(err, models.ErrMaxFileSize) {
					t.mainPage("file size exceeded")
					return
				}
				if err != nil {
					t.mainPage("Unknown error")
					return
				}

				t.mainPage("OK")
			}).
			AddButton("Back", func() {
				t.mainPage("")
			})

		frame := tview.NewFrame(form).SetBorders(0, 0, 0, 1, 4, 4).
			AddText("TAB - for switching between fields | Enter - for select", false, tview.AlignLeft, tcell.ColorWhite).
			AddText(message, false, tview.AlignRight, tcell.ColorRed)

		t.pages.AddPage("file", frame, true, true)
		t.pages.SwitchToPage("file")

	case models.Delete:
		form := tview.NewForm().
			AddInputField("Record ID:", "", 30, nil, func(id string) {
				res, err := strconv.Atoi(id)
				if err != nil {
					t.filePage("invalid id")
					return
				}
				t.client.Record.RecordID = int32(res)
			}).
			AddButton("Delete", func() {

				err := t.client.DeleteRecordByID()

				if errors.Is(err, status.Error(codes.InvalidArgument, "")) {
					t.mainPage("InvalidArgument")
					return
				}
				if err != nil {
					t.mainPage("Unknown error")
					return
				}

				t.mainPage("OK")
			}).
			AddButton("Back", func() {
				t.mainPage("")
			})

		frame := tview.NewFrame(form).SetBorders(0, 0, 0, 1, 4, 4).
			AddText("TAB - for switching between fields | Enter - for select", false, tview.AlignLeft, tcell.ColorWhite).
			AddText(message, false, tview.AlignRight, tcell.ColorRed)

		t.pages.AddPage("file", frame, true, true)
		t.pages.SwitchToPage("file")
	}
}
