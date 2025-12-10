package auth

import (
	"log"
	"net/http"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func registerWindow(w fyne.Window) fyne.CanvasObject {
	w.SetTitle("Registration")
	w.Resize(fyne.NewSize(300, 400))

	labelLogin, loginField := createInputField("Insert login", false)
	labelPassword, passwordField := createInputField("Insert password", true)
	labelCheckPassword, passwordCheckField := createInputField("Insert password again", true)

	registerButton := widget.NewButton("Registration", func() {
		if len(strings.TrimSpace(loginField.Text)) == 0 || len(strings.TrimSpace(passwordField.Text)) == 0 {
			return
		}

		if isOnlyDigit(loginField.Text) {
			dialog := customErrorMsg("Login cannot consist only of numbers", w)
			dialog.Show()
			return
		}

		if passwordCheckField.Text != passwordField.Text {
			dialog := customErrorMsg("Passwords don't match", w)
			dialog.Show()
			return
		}

		if !checkPassword(passwordField.Text) {
			dialog := customErrorMsg(`Password must be at least 8 characters long and contain numbers and letters of both upper and lower case`, w)
			dialog.Show()
			return
		}

		resp, err := authApi(loginField.Text, passwordCheckField.Text, true)
		if err != nil || resp.HttpCode != http.StatusCreated {
			log.Println(err)
			dialog := customErrorMsg("Try later", w)
			dialog.Show()
			return
		}

		w.SetContent(LoginWindow(w))
	})

	backwardButton := widget.NewButton("Back", func() {
		w.SetContent(LoginWindow(w))
	})

	content := container.NewVBox(labelLogin, loginField, labelPassword, passwordField, labelCheckPassword, passwordCheckField, registerButton, backwardButton)
	return content
}
