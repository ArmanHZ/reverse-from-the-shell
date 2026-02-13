package ui

import (
	"bytes"
	"encoding/base64"
	"html/template"
	"net/url"
	"rvfs/data"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (a *App) focusNext() {
	if len(a.focusables) == 0 {
		return
	}

	a.focusIndex = (a.focusIndex + 1) % len(a.focusables)
	a.app.SetFocus(a.focusables[a.focusIndex])
}

func (a *App) focusPrev() {
	if len(a.focusables) == 0 {
		return
	}

	a.focusIndex--
	if a.focusIndex < 0 {
		a.focusIndex = len(a.focusables) - 1
	}

	a.app.SetFocus(a.focusables[a.focusIndex])
}

// TODO: Maybe a better name
func (a *App) initInputCapture() {
	a.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {

		case tcell.KeyTab:
			a.focusNext()
			return nil

		case tcell.KeyBacktab:
			a.focusPrev()
			return nil
		}

		return event
	})
}

func (a *App) triggerGlobalUiUpdate() {
	ip := a.ipField.GetText()
	port := a.portField.GetText()
	_, shell := a.shellTypeSelect.GetCurrentOption()

	payload := data.ReverseShellCommands[a.payloadTableRow].Command

	sDec, err := base64.StdEncoding.DecodeString(payload)
	if err != nil {
		panic(err)
	}

	tmpl, err := template.New("").Parse(string(sDec))
	if err != nil {
		panic(err)
	}

	var _, currentEncodingType = a.encodingTypeSelect.GetCurrentOption()
	if currentEncodingType == "None" {
		// Adding colors
		shell = "[blue]" + shell + "[white]"
		port = "[green]" + port + "[white]"
		ip = "[yellow]" + ip + "[white]"
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, map[string]string{"Shell": shell, "Port": port, "Ip": ip})
	var bufferAsString string = buf.String()

	if err != nil {
		panic(err)
	}

	// FIXME: Base64 url safe removes important characters like ">"...
	switch currentEncodingType {
	case "Base64":
		bufferAsString = base64.StdEncoding.EncodeToString([]byte(bufferAsString))
	case "Base64UrlSafe":
		bufferAsString = base64.URLEncoding.EncodeToString([]byte(bufferAsString))
	case "UrlEncoding":
		bufferAsString = url.QueryEscape(bufferAsString)
	case "UrlAndBase64":
		bufferAsString = base64.StdEncoding.EncodeToString([]byte(url.QueryEscape(bufferAsString)))
	case "UrlAndBase64UrlSafe":
		bufferAsString = base64.URLEncoding.EncodeToString([]byte(url.QueryEscape(bufferAsString)))
	}

	a.shellPayloadDisplay.SetText(bufferAsString)
}

func (a *App) initIpFieldEvents() {
	a.ipField.SetChangedFunc(func(text string) {
		a.triggerGlobalUiUpdate()
	})
}

func (a *App) initPortFieldEvents() {
	a.portField.SetChangedFunc(func(text string) {
		// FIXME: Clean this up and rename the variables.
		tmp := strings.Fields(a.listenerCommand.GetText(true))
		var tmp2 string
		if len(a.portField.GetText()) > 0 {
			tmp2 = strings.Join(tmp[:len(tmp)-1], " ")
			a.listenerCommand.SetText(tmp2 + " " + text)
		}

		a.triggerGlobalUiUpdate()
	})
}

func (a *App) initListenerTypeSelectEvents() {
	var dropDownOptions []string
	for _, v := range data.Listeners {
		dropDownOptions = append(dropDownOptions, v.Name)
	}

	a.listenerTypeSelect.SetOptions(dropDownOptions, func(t string, i int) {
		tmpl, err := template.New("listener").Parse(data.Listeners[i].Payload)
		if err != nil {
			panic(err)
		}

		var buf bytes.Buffer
		tmpl.Execute(&buf, map[string]string{"Port": a.portField.GetText(), "Ip": a.ipField.GetText()})

		a.listenerCommand.SetText(buf.String())
	}).
		SetCurrentOption(0)
}

func (a *App) initTargetOsTypeSelectEvents() {
	a.targetOsTypeSelect.SetOptions(data.OSTypes, nil).
		SetCurrentOption(0)
}

// FIXME: This name disregards other payload types.
func (a *App) initReverseShellTableEvents() {
	// XXX: Should I do this part in this file? Who knows...
	for row, text := range data.ReverseShellCommands {
		cell := tview.NewTableCell(text.Name).
			SetAlign(tview.AlignLeft)

		a.shellCommandTable.SetCell(row, 0, cell)
	}

	a.shellCommandTable.SetBlurFunc(func() {
		a.shellCommandTable.SetSelectable(false, false)
	})

	a.shellCommandTable.SetFocusFunc(func() {
		a.shellCommandTable.SetSelectable(true, false)
	})

	a.shellCommandTable.SetSelectionChangedFunc(func(row, column int) {
		a.payloadTableRow = row
		a.payloadTableColumn = 0

		a.triggerGlobalUiUpdate()
	})

	a.shellCommandTable.Select(0, 0)
}

func (a *App) initShellPayloadSelectEvents() {
	a.shellTypeSelect.SetSelectedFunc(func(text string, index int) {
		a.triggerGlobalUiUpdate()
	})
}

func (a *App) initClipboardEvents() {
	a.listenerCopyButton.SetSelectedFunc(func() {
		a.CopyToClipBoard(0)
	})

	a.shellCommandTable.SetSelectedFunc(func(row, column int) {
		a.CopyToClipBoard(1)
	})

	a.shellPayloadCopyButton.SetSelectedFunc(func() {
		a.CopyToClipBoard(1)
	})
}

// TODO: Maybe better names?
func (a *App) bindEvents() {
	a.initInputCapture()
	a.initIpFieldEvents()
	a.initPortFieldEvents()
	a.initListenerTypeSelectEvents()
	a.initTargetOsTypeSelectEvents()
	a.initReverseShellTableEvents()
	a.initShellPayloadSelectEvents()
	a.initClipboardEvents()
}
