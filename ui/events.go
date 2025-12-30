package ui

import (
	"bytes"
	"html/template"
	"rvfs/data"
)

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
