package ui

import (
	"bytes"
	"encoding/base64"
	"html/template"
	"rvfs/data"

	"github.com/rivo/tview"
)

func Spacer() *tview.Box {
	return tview.NewBox()
}

func (a *App) registerFocusable(p tview.Primitive) {
	a.focusables = append(a.focusables, p)
}

func (a *App) triggerGlobalUiUpdate() {
	ip := a.ipField.GetText()
	port := a.portField.GetText()
	// shell type

	payload := data.ReverseShellCommands[a.payloadTableRow].Command

	sDec, _ := base64.StdEncoding.DecodeString(payload)

	tmpl, err := template.New("").Parse(string(sDec))
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	// TODO: I put bash here temporarily.
	tmpl.Execute(&buf, map[string]string{"Shell": "bash", "Port": port, "Ip": ip})

	a.reverseShellCommandDisplay.SetText(buf.String())
}
