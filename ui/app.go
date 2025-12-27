package ui

import "github.com/rivo/tview"

type App struct {
	app *tview.Application

	// UI state
	mainGrid *tview.Grid

	// Header widgets
	ipField    *tview.InputField
	portField  *tview.InputField
	typeSelect *tview.DropDown
	cmdView    *tview.TextView
}

func New() *App {
	a := &App{
		app: tview.NewApplication(),
	}

	a.buildUI()
	// a.bindEvents()

	return a
}

func (a *App) Run() error {
	return a.app.Run()
}
