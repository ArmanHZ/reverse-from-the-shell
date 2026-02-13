package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type App struct {
	app *tview.Application

	clipboardError error

	focusables []tview.Primitive
	focusIndex int

	// UI elements below this line
	mainGrid *tview.Grid

	// Header widgets
	ipField            *tview.InputField
	portField          *tview.InputField
	listenerCommand    *tview.TextView
	listenerTypeSelect *tview.DropDown
	listenerCopyButton *tview.Button

	// Main content widgets
	targetOsTypeSelect     *tview.DropDown
	shellCommandTable      *tview.Table
	shellPayloadDisplay    *tview.TextView
	shellTypeSelect        *tview.DropDown
	shellPayloadCopyButton *tview.Button
	encodingTypeSelect     *tview.DropDown

	// Tab buttons
	// TODO: Maybe do something like an enum
	// tabs -> 0: Reverse, 1: Bind, 2: MSFVenom, 3: HoaxShell
	tabButtons []*tview.Button

	payloadTableRow    int
	payloadTableColumn int // This is always 0 anyways. But prevents magic num.
}

/* Why is this done?
 * Some term themes, such as Kanagawa-Wave, use bright colors for core tcell
 * colors such as blue. This makes the UI very dificult to read and some
 * components just look ugly.
 * If you are having that exact issue, I have set this "dark mode" theme as
 * a bandaid solution.
 * This will only be in use when the user provides the `-color-fix` flag.
 */
var darkModeColorFix = tview.Theme{
	PrimitiveBackgroundColor:    tcell.ColorBlack,
	ContrastBackgroundColor:     tcell.ColorDarkBlue,
	MoreContrastBackgroundColor: tcell.ColorDarkGreen,
	BorderColor:                 tcell.ColorGray,
	TitleColor:                  tcell.ColorGray,
	GraphicsColor:               tcell.ColorGray,
	PrimaryTextColor:            tcell.ColorGray,
	SecondaryTextColor:          tcell.ColorDarkGoldenrod,
	TertiaryTextColor:           tcell.ColorYellow,
	InverseTextColor:            tcell.ColorBlack,
	ContrastSecondaryTextColor:  tcell.ColorDarkGray,
}

func New(isColorFix bool, userIp string) *App {
	if isColorFix {
		tview.Styles = darkModeColorFix
	}

	a := &App{
		app: tview.NewApplication(),
	}

	a.buildUI()
	a.bindEvents()
	a.initInputCapture()

	// TODO: Idk if setting up the cmdline variables here is good practice.
	// but also don't want to complicate things with passing variables around.
	if userIp != "" {
		a.ipField.SetText(userIp)
	}

	a.app.SetFocus(a.focusables[0])

	// FIXME: This messes up with the tab movement.
	a.app.EnableMouse(true) // For testing different widgets. Might disable later or let it be.

	return a
}

func (a *App) Run() error {
	return a.app.Run()
}
