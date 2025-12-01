package main

import (
	"voiceChatClient/auth"

	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("VoiceChat")

	w.SetContent(auth.LoginWindow(w))
	w.ShowAndRun()
}
