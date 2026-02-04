package ui

import (
	"bytes"
	"encoding/base64"
	"html/template"

	//	"regexp"
	"rvfs/data"

	"github.com/rivo/tview"
	"golang.design/x/clipboard"
)

// TODO: Will use these later.
// idk... I copied and pasted these from stack overflow lol
// var IpFieldCheck = regexp.MustCompile(`^(((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.|$)){4})`)
// var PortFieldCheck = regexp.MustCompile(`^([1-9][0-9]{0,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5])$`)

func Spacer() *tview.Box {
	return tview.NewBox()
}

func (a *App) registerFocusable(p tview.Primitive) {
	a.focusables = append(a.focusables, p)
}

// XXX: This function may need to be put in the `events.go` file.
func (a *App) triggerGlobalUiUpdate() {
	ip := a.ipField.GetText()
	port := a.portField.GetText()
	_, shell := a.shellPayloadSelect.GetCurrentOption()

	payload := data.ReverseShellCommands[a.payloadTableRow].Command

	sDec, err := base64.StdEncoding.DecodeString(payload)
	if err != nil {
		panic(err)
	}

	tmpl, err := template.New("").Parse(string(sDec))
	if err != nil {
		panic(err)
	}

	// Adding colors
	shell = "[blue]" + shell + "[white]"
	port = "[green]" + port + "[white]"
	ip = "[yellow]" + ip + "[white]"

	var buf bytes.Buffer
	tmpl.Execute(&buf, map[string]string{"Shell": shell, "Port": port, "Ip": ip})

	a.reverseShellCommandDisplay.SetText(buf.String())
}

// TODO: Maybe use something other than index. Enums maybe, idk.
func (a *App) CopyToClipBoard(index int) {
	a.clipboardError = clipboard.Init()
	if a.clipboardError != nil {
		panic(a.clipboardError)
	}

	switch index {
	case 0: // Listener copy button
		clipboard.Write(clipboard.FmtText, []byte(a.listenerCommand.GetText(true)))

	case 1: // Payload copy button
		clipboard.Write(clipboard.FmtText, []byte(a.reverseShellCommandDisplay.GetText(true)))
	}
}
