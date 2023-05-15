package tui

import (
	"fmt"

	"github.com/RomanIkonnikov93/keeper/client/internal/models"
	"github.com/rivo/tview"
)

func (t *TUI) outputPage(message string) {

	switch t.client.Record.RecordType {
	case models.Credentials:

		list := tview.NewList().
			AddItem(fmt.Sprint(t.client.Record.RecordID), "record ID", 'i', nil).
			AddItem(t.client.Record.Description, "description", 'd', nil).
			AddItem(t.client.Record.Metadata, "metadata", 'm', nil).
			AddItem(t.client.Record.Login, "login", 'l', nil).
			AddItem(t.client.Record.Password, "password", 'p', nil).
			AddItem("", "Back", 'b', func() {
				t.mainPage("")
			}).
			AddItem("", "Quit", 'q', func() {
				t.Application.Stop()
			})

		t.pages.AddPage("output", list, true, true)
		t.pages.SwitchToPage("output")

		t.client.CleanRecordFields()

	case models.Card:

		list := tview.NewList().
			AddItem(fmt.Sprint(t.client.Record.RecordID), "record ID", 'i', nil).
			AddItem(t.client.Record.Description, "description", 'd', nil).
			AddItem(t.client.Record.Metadata, "metadata", 'm', nil).
			AddItem(t.client.Record.Card, "card", 'l', nil).
			AddItem("", "Back", 'b', func() {
				t.mainPage("")
			}).
			AddItem("", "Quit", 'q', func() {
				t.Application.Stop()
			})

		t.pages.AddPage("output", list, true, true)
		t.pages.SwitchToPage("output")

		t.client.CleanRecordFields()

	case models.File:

	}
}
