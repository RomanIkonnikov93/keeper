package tui

import (
	"github.com/RomanIkonnikov93/keeper/client/internal/models"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (t *TUI) mainPage(message string) {

	form := tview.NewForm()

	form.AddDropDown("Select an option (hit Enter): ", []string{"credentials", "card", "file"}, -1, func(option string, optionIndex int) {
		switch option {
		case "credentials":
			t.client.Record.RecordType = models.Credentials
		case "card":
			t.client.Record.RecordType = models.Card
		case "file":
			t.client.Record.RecordType = models.File
		}
	}).AddButton("Add record", func() {

		if t.client.Record.RecordType == "" {
			t.mainPage("record type not selected")
			return
		}

		t.client.Record.ActionType = models.Add

		switch t.client.Record.RecordType {
		case models.Credentials:
			t.credentialsPage("")
		case models.Card:
			t.cardPage("")
		case models.File:
			t.filePage("")
		}

	}).AddButton("Get record", func() {

		if t.client.Record.RecordType == "" {
			t.mainPage("record type not selected")
			return
		}

		t.client.Record.ActionType = models.Get

		switch t.client.Record.RecordType {
		case models.Credentials:
			t.credentialsPage("")
		case models.Card:
			t.cardPage("")
		case models.File:
			t.filePage("")
		}
	}).AddButton("Update record", func() {

		if t.client.Record.RecordType == "" {
			t.mainPage("record type not selected")
			return
		}

		t.client.Record.ActionType = models.Update

		switch t.client.Record.RecordType {
		case models.Credentials:
			t.credentialsPage("")
		case models.Card:
			t.cardPage("")
		case models.File:
			t.filePage("")
		}
	}).AddButton("Delete record", func() {

		if t.client.Record.RecordType == "" {
			t.mainPage("record type not selected")
			return
		}

		t.client.Record.ActionType = models.Delete

		switch t.client.Record.RecordType {
		case models.Credentials:
			t.credentialsPage("")
		case models.Card:
			t.cardPage("")
		case models.File:
			t.filePage("")
		}

	}).AddButton("Get all records", func() {

		if t.client.Record.RecordType == "" {
			t.mainPage("record type not selected")
			return
		}

		t.client.Record.ActionType = models.GetAll

		t.outputAllPage("")

	}).SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 113 {
			t.Application.Stop()
		}
		return event
	})

	frame := tview.NewFrame(form).SetBorders(0, 0, 0, 3, 4, 4).
		AddText("TAB - for switching between fields | Enter - for select | (q) - to quit", false, tview.AlignLeft, tcell.ColorWhite).
		AddText(message, false, tview.AlignRight, tcell.ColorRed)

	t.pages.AddPage("main", frame, true, true)
	t.pages.SwitchToPage("main")
}
