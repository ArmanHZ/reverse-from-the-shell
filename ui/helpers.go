package ui

import (

	//	"regexp"

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
		clipboard.Write(clipboard.FmtText, []byte(a.shellPayloadDisplay.GetText(true)))
	}
}
