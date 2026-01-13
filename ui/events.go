package ui

import (
	"bytes"
	"html/template"
	"rvfs/data"

	"github.com/gdamore/tcell/v2"
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

// TODO: Split each component's events to their own respective functions.
func (a *App) bindEvents() {
	a.portField.SetChangedFunc(func(text string) {
		a.listenerCommand.SetText("nc -lvnp " + text)
	})

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
	})
}
