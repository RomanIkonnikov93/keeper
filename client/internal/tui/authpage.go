package tui

import (
	"errors"

	"github.com/RomanIkonnikov93/keeper/client/internal/models"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (t *TUI) authPage(message string) {

	form := tview.NewForm().
		AddTextView("Keeper", "This application allows you to safely store private information on a remote server.\nSuch as: passwords, payment card data, files.", 100, 2, true, false).
		AddInputField("Username:", "", 20, nil, func(username string) {
			t.client.Auth.Login = username
		}).
		AddPasswordField("Password:", "", 20, '*', func(password string) {
			t.client.Auth.Password = password
		}).
		AddButton("Sign Up", func() {

			err := t.client.RegistrationUser()

			if errors.Is(err, status.Error(codes.InvalidArgument, "")) {
				t.authPage("InvalidArgument")
				return
			}
			if errors.Is(err, status.Error(codes.AlreadyExists, "")) {
				t.authPage("AlreadyExists")
				return
			}
			if err != nil {
				t.authPage("Unknown error")
				return
			}

			t.mainPage("")
		}).
		AddButton("Log In", func() {

			err := t.client.LoginUser()

			if errors.Is(err, status.Error(codes.InvalidArgument, "")) {
				t.authPage("InvalidArgument")
				return
			}
			if errors.Is(err, status.Error(codes.AlreadyExists, "")) {
				t.authPage("AlreadyExists")
				return
			}
			if err != nil {
				t.authPage("Unknown error")
				return
			}

			t.mainPage("")
		}).
		AddButton("Quit", func() {

			t.Application.Stop()

		}).
		AddTextView("Build Version:", models.BuildVersion, 0, 1, true, false)

	frame := tview.NewFrame(form).SetBorders(0, 0, 0, 1, 4, 4).
		AddText("TAB - for switching between fields | Enter - for select", false, tview.AlignLeft, tcell.ColorWhite).
		AddText(message, false, tview.AlignRight, tcell.ColorRed)

	t.pages.AddPage("auth", frame, true, true)
	t.pages.SwitchToPage("auth")
}
