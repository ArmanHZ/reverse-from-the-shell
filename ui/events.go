package ui

import (
	"bytes"
	"encoding/base64"
	"html/template"
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

// FIXME: Table content also needs to change.
func (a *App) initPortFieldEvents() {
	a.portField.SetChangedFunc(func(text string) {
		// FIXME: Clean this up and rename the variables.
		tmp := strings.Fields(a.listenerCommand.GetText(true))
		var tmp2 string
		if len(a.portField.GetText()) > 0 {
			tmp2 = strings.Join(tmp[:len(tmp)-1], " ")
			a.listenerCommand.SetText(tmp2 + " " + text)
		}
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

// FIXME: When the user changes IP and/or Port, the change does not reflect
// onto the table until you interract with the table again.
func (a *App) initReverseShellTableEvents() {
	// XXX: Should I do this part in this file? Who knows...
	for row, text := range data.ReverseShellCommands {
		cell := tview.NewTableCell(text.Name).
			SetAlign(tview.AlignLeft)

		a.reverseShellSelect.SetCell(row, 0, cell)
	}

	a.reverseShellSelect.SetSelectionChangedFunc(func(row, column int) {
		tmp := data.ReverseShellCommands[row].Command

		sDec, _ := base64.StdEncoding.DecodeString(tmp)

		tmpl, err := template.New("").Parse(string(sDec))
		if err != nil {
			panic(err)
		}

		var buf bytes.Buffer
		// TODO: I put bash here temporarily.
		tmpl.Execute(&buf, map[string]string{"Shell": "bash", "Port": a.portField.GetText(), "Ip": a.ipField.GetText()})

		a.reverseShellCommandDisplay.SetText(buf.String())
	})

	a.reverseShellSelect.Select(0, 0)
}

// TODO: Maybe better names?
func (a *App) bindEvents() {
	a.initInputCapture()
	a.initPortFieldEvents()
	a.initListenerTypeSelectEvents()
	a.initTargetOsTypeSelectEvents()
	a.initReverseShellTableEvents()
}
