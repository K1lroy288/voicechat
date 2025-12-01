package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"unicode"
	"voiceChatClient/config"
	"voiceChatClient/model"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func authApi(username, password string, register bool) (*model.ServerResponse, error) {
	cfg := config.GetConfig()

	url := cfg.HttpTls + "://" + cfg.ServerHost + ":" + cfg.ServerPort
	if register {
		url += "/auth/register"
	} else {
		url += "/auth/login"
	}

	user := model.User{
		Username: username,
		Password: password,
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		return nil, fmt.Errorf("register JSON marshal error: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("register request create error: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("register request send error: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("register response read error: %w", err)
	}

	var serverResponse model.ServerResponse
	if err := json.Unmarshal(body, &serverResponse); err != nil {
		return nil, fmt.Errorf("register response decode error: %w", err)
	}

	serverResponse.HttpCode = resp.StatusCode

	return &serverResponse, nil
}

func customErrorMsg(desc string, w fyne.Window) dialog.Dialog {
	content := widget.NewLabel(desc)
	content.Wrapping = fyne.TextWrapWord
	scrollableContent := container.NewScroll(content)

	dialog := dialog.NewCustom("Login or Register error", "OK", scrollableContent, w)
	dialog.Resize(fyne.NewSize(300, 200))

	return dialog
}

func createInputField(labelText string, isPassword bool) (*widget.Label, *widget.Entry) {
	label := widget.NewLabel(labelText)
	if isPassword {
		return label, widget.NewPasswordEntry()
	}

	return label, widget.NewEntry()
}

func checkPassword(password string) bool {
	var lets int
	for _, char := range password {
		if unicode.IsLetter(char) {
			lets += 1
		}
	}

	if lets == len(password) || len(password) < 8 || lets == 0 {
		return false
	}

	if password == strings.ToLower(password) {
		return false
	}

	return true
}

func isOnlyDigit(str string) bool {
	for _, c := range str {
		if !unicode.IsDigit(c) {
			return false
		}

	}

	return true
}
