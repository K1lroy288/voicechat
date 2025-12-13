package auth

import (
	"log"
	"strings"
	"voiceChatClient/config"
	mainwindow "voiceChatClient/main_window"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func LoginWindow(w fyne.Window) fyne.CanvasObject {
	w.SetTitle("Login")
	w.Resize(fyne.NewSize(300, 400))

	labelLogin, loginField := createInputField("Insert login", false)
	labelPassword, passwordField := createInputField("Insert password", true)

	loginButton := widget.NewButton("Login", func() {
		if len(strings.TrimSpace(loginField.Text)) == 0 || len(strings.TrimSpace(passwordField.Text)) == 0 {
			return
		}

		log.Println(loginField.Text, "-", passwordField.Text, "-")

		_, _ = authApi("1", "1", true)

		resp, err := authApi(loginField.Text, passwordField.Text, false)
		if (err != nil) || ((resp.HttpCode != 200) && (resp.HttpCode != 500)) {
			log.Println(err)
			dialog := customErrorMsg("Try later", w)
			dialog.Show()
			return
		}

		log.Println(resp.Error)

		if resp.HttpCode == 500 {
			log.Println(err)
			dialog := customErrorMsg("Wrong username or password", w)
			dialog.Show()
			return
		}

		config.UserData.Token = resp.Token
		config.UserData.Username = resp.User.Username
		config.UserData.UserId = resp.User.Id
		w.SetContent(mainwindow.VoiceChatMainWindow(w))
	})

	registerButton := widget.NewButton("Register", func() {
		w.SetContent(registerWindow(w))
	})

	content := container.NewVBox(labelLogin, loginField, labelPassword, passwordField, loginButton, registerButton)
	return content
}
