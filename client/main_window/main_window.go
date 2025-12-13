package mainwindow

import (
	"strconv"
	"voiceChatClient/config"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func VoiceChatMainWindow(w fyne.Window) fyne.CanvasObject {
	w.SetTitle(config.UserData.Username + "  #" + strconv.Itoa(config.UserData.UserId))
	w.Resize(fyne.NewSize(300, 400))

	friendList := widget.NewList(
		func() int {
			return 5
		},
		func() fyne.CanvasObject {
			return widget.NewButton("", func() {

			})
		},
		func(id int, item fyne.CanvasObject) {
			item.(*widget.Button).SetText("123")
		},
	)

	friendSearchButton := widget.NewButton("Find Friend", func() {

	})

	friendRequests := widget.NewButton("Friend Requests", func() {

	})

	content := container.NewBorder(
		container.NewHBox(friendSearchButton, friendRequests),
		nil,
		nil,
		nil,
		friendList,
	)

	return content
}
