package ui

import "github.com/rivo/tview"

func Spacer() *tview.Box {
	return tview.NewBox()
}

func (a *App) registerFocusable(p tview.Primitive) {
	a.focusables = append(a.focusables, p)
}
