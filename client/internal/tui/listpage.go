package tui

import (
	"fmt"

	"github.com/RomanIkonnikov93/keeper/client/internal/models"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (t *TUI) outputAllPage(message string) {

	switch t.client.Record.RecordType {
	case models.Credentials:

		table := tview.NewTable().
			SetBorders(true)

		arr := make([]string, 0)
		arr = append(arr, "Record ID", "Description", "Metadata", "Login", "Password")

		t.client.Mutex.Lock()
		for _, r := range t.client.Store.Credentials {
			arr = append(arr, fmt.Sprint(r.RecordID))
			arr = append(arr, r.Description)
			arr = append(arr, r.Metadata)
			arr = append(arr, r.Login)
			arr = append(arr, r.Password)
		}
		t.client.Mutex.Unlock()

		cols, rows := 5, len(t.client.Store.Credentials)+1
		word := 0
		for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {
				color := tcell.ColorWhite
				if c < 1 || r < 1 {
					color = tcell.ColorYellow
				}
				table.SetCell(r, c,
					tview.NewTableCell(arr[word]).
						SetTextColor(color).
						SetAlign(tview.AlignCenter))
				word = (word + 1) % len(arr)
			}
		}
		table.Select(0, 0).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
			if key == tcell.KeyEscape {
				t.Application.Stop()
			}
			if key == tcell.KeyEnter {
				table.SetSelectable(true, true)
			}
		}).SetSelectedFunc(func(row int, column int) {
			table.GetCell(row, column).SetTextColor(tcell.ColorRed)
			table.SetSelectable(false, false)
		}).
			SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
				if event.Key() == tcell.KeyBackspace {
					t.mainPage("")
				}
				return event
			})

		frame := tview.NewFrame(table).SetBorders(0, 0, 0, 3, 4, 4).
			AddText("Enter - for select | use: ⇅ ⇄  to to navigate | (Backspace ←) - back", false, tview.AlignLeft, tcell.ColorWhite).
			AddText(message, false, tview.AlignRight, tcell.ColorRed)

		t.pages.AddPage("outputAll", frame, true, true)
		t.pages.SwitchToPage("outputAll")

	case models.Card:

		table := tview.NewTable().
			SetBorders(true)

		arr := make([]string, 0)
		arr = append(arr, "Record ID", "Description", "Metadata", "Card")

		t.client.Mutex.Lock()
		for _, r := range t.client.Store.Cards {
			arr = append(arr, fmt.Sprint(r.RecordID))
			arr = append(arr, r.Description)
			arr = append(arr, r.Metadata)
			arr = append(arr, r.Card)
		}
		t.client.Mutex.Unlock()

		cols, rows := 4, len(t.client.Store.Cards)+1
		word := 0
		for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {
				color := tcell.ColorWhite
				if c < 1 || r < 1 {
					color = tcell.ColorYellow
				}
				table.SetCell(r, c,
					tview.NewTableCell(arr[word]).
						SetTextColor(color).
						SetAlign(tview.AlignCenter))
				word = (word + 1) % len(arr)
			}
		}
		table.Select(0, 0).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
			if key == tcell.KeyEscape {
				t.Application.Stop()
			}
			if key == tcell.KeyEnter {
				table.SetSelectable(true, true)
			}
		}).SetSelectedFunc(func(row int, column int) {
			table.GetCell(row, column).SetTextColor(tcell.ColorRed)
			table.SetSelectable(false, false)
		}).
			SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
				if event.Key() == tcell.KeyBackspace {
					t.mainPage("")
				}
				return event
			})

		frame := tview.NewFrame(table).SetBorders(0, 0, 0, 3, 4, 4).
			AddText("Enter - for select | use: ⇅ ⇄  to to navigate | (Backspace ←) - back", false, tview.AlignLeft, tcell.ColorWhite).
			AddText(message, false, tview.AlignRight, tcell.ColorRed)

		t.pages.AddPage("outputAll", frame, true, true)
		t.pages.SwitchToPage("outputAll")

	case models.File:

	}
}
