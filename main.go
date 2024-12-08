package main

import (
	"context"
	"time"

	"github.com/gonutz/wui/v2"
)

func main() {
	workingStatus := true
	reauthguiFont, _ := wui.NewFont(wui.FontDesc{
		Name:   "Segoe UI",
		Height: -11,
	})

	reauthgui := wui.NewWindow()
	reauthgui.SetFont(reauthguiFont)
	reauthgui.SetInnerSize(385, 291)
	reauthgui.SetTitle("wifi-reauth-go")
	reauthgui.SetHasMaxButton(false)
	reauthgui.SetResizable(false)
	reauthgui.SetState(wui.WindowMinimized)

	button1 := wui.NewButton()
	button1.SetBounds(10, 110, 120, 24)
	button1.SetText("Stop/start daemon")
	reauthgui.Add(button1)

	button2 := wui.NewButton()
	button2.SetBounds(10, 85, 120, 24)
	button2.SetText("Logout")
	reauthgui.Add(button2)

	daemonStatusLabel := wui.NewLabel()
	daemonStatusLabel.SetBounds(15, 25, 50, 13)
	daemonStatusLabel.SetText("Daemon:")
	reauthgui.Add(daemonStatusLabel)

	daemonStatusText := wui.NewLabel()
	daemonStatusText.SetBounds(65, 25, 150, 13)
	daemonStatusText.SetText("Running")
	reauthgui.Add(daemonStatusText)

	lastLoginTime := wui.NewLabel()
	lastLoginTime.SetBounds(15, 45, 150, 13)
	lastLoginTime.SetText("Last login:")
	reauthgui.Add(lastLoginTime)

	label1 := wui.NewLabel()
	label1.SetBounds(70, 45, 180, 14)
	label1.SetText("00:00:00 01/01/1970")
	reauthgui.Add(label1)

	label2 := wui.NewLabel()
	label2.SetBounds(145, 114, 300, 29)
	label2.SetText("https://github.com/michioxd/wifi-reauth-go")
	reauthgui.Add(label2)

	label3 := wui.NewLabel()
	label3.SetBounds(15, 65, 150, 13)
	label3.SetText("Status:")
	reauthgui.Add(label3)

	statusText := wui.NewLabel()
	statusText.SetBounds(55, 65, 150, 13)
	statusText.SetText("Idle")
	reauthgui.Add(statusText)

	labelbruh := wui.NewLabel()
	labelbruh.SetBounds(145, 110, 300, 13)
	labelbruh.SetText("For Free Wi-MESH - KTX TĐHBK - ĐHĐN")
	reauthgui.Add(labelbruh)

	loggingBox := wui.NewTextEdit()
	loggingBox.SetEnabled(true)
	loggingBox.SetBounds(10, 145, 365, 137)
	reauthgui.Add(loggingBox)

	button1.SetOnClick(func() {
		if workingStatus {
			daemonStatusText.SetText("Stopped")
			statusText.SetText("Idle")
			workingStatus = false
		} else {
			daemonStatusText.SetText("Running")
			statusText.SetText("Checking...")
			workingStatus = true
		}
	})

	button2.SetOnClick(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		Logout(ctx)

		statusText.SetText("Logged out")
		loggingBox.SetText(loggingBox.Text() + "Logout: " + time.Now().Format("15:04:05 02/01/2006") + "\r\n")
	})

	go func() {
		for {
			if workingStatus {
				checkAndLogin(statusText, loggingBox, label1)
			}
			time.Sleep(1 * time.Second)
		}
	}()

	reauthgui.Show()
}

func checkAndLogin(statusText *wui.Label, loggingBox *wui.TextEdit, label1 *wui.Label) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if !CheckAuth(ctx) {
		statusText.SetText("Logging in...")
		res, err := Login(ctx)
		if err != "" {
			statusText.SetText("Failed, retrying...")
			loggingBox.SetText(loggingBox.Text() + "Login failed: " + time.Now().Format("15:04:05 02/01/2006") + " - " + err + "\r\n")
		} else if res {
			statusText.SetText("OK")
			loggingBox.SetText(loggingBox.Text() + "Login success: " + time.Now().Format("15:04:05 02/01/2006") + "\r\n")
			label1.SetText(time.Now().Format("15:04:05 02/01/2006"))
		}
	} else {
		statusText.SetText("Idle")
	}
}
