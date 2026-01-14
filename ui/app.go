package ui

import "github.com/rivo/tview"

type App struct {
	app *tview.Application

	focusables []tview.Primitive
	focusIndex int

	// UI elements below this line
	mainGrid *tview.Grid

	// Header widgets
	ipField            *tview.InputField
	portField          *tview.InputField
	listenerCommand    *tview.TextView
	listenerTypeSelect *tview.DropDown // FIXME: I don't like this name.
	listenerCopyButton *tview.Button

	// Main content widgets
	targetOsTypeSelect         *tview.DropDown
	reverseShellSelect         *tview.Table
	reverseShellCommandDisplay *tview.TextView

	// Tab buttons
	// TODO: Maybe do something like an enum
	// tabs -> 0: Reverse, 1: Bind, 2: MSFVenom, 3: HoaxShell
	tabButtons []*tview.Button
}

func New() *App {
	a := &App{
		app: tview.NewApplication(),
	}

	a.buildUI()
	a.bindEvents()
	a.initInputCapture()

	a.app.SetFocus(a.focusables[0])

	// FIXME: This messes up with the tab movement.
	a.app.EnableMouse(true) // For testing different widgets. Might disable later or let it be.

	return a
}

func (a *App) Run() error {
	return a.app.Run()
}
