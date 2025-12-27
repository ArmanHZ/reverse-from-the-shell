package ui

import (
	// "github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func buildHeader() {

}

func (a *App) buildUI() {
	a.mainGrid = tview.NewGrid().
		SetRows(1, 4, 1, 0).
		SetColumns(0).
		SetBorders(true)

	title := tview.NewTextView().
		SetDynamicColors(true).
		SetText("Reverse Shell Generator").
		SetTextAlign(tview.AlignCenter)

	headerGrid := tview.NewGrid().
		SetRows(0).
		SetColumns(0, 4, 0)

	ipField := tview.NewInputField().
		SetLabel("IP: ").
		SetText("10.10.10.10")

	portField := tview.NewInputField().
		SetLabel("Port: ").
		SetText("9001")

	ipPortFlex := tview.NewFlex().
		AddItem(tview.NewTextView().SetText("IP & Port"), 0, 1, false).
		SetDirection(tview.FlexRow).
		AddItem(Spacer(), 0, 1, false).
		AddItem(tview.NewFlex().
			AddItem(ipField, 0, 4, true).
			AddItem(Spacer(), 0, 1, false).
			AddItem(portField, 0, 4, true), 0, 1, true)

	listenerCommand := tview.NewTextView().
		SetText("nc -lvnp " + portField.GetText()).
		SetTextAlign(tview.AlignCenter)

	portField.SetChangedFunc(func(text string) {
		listenerCommand.SetText("nc -lvnp " + text)
	})

	listenerTypeDropdown := tview.NewDropDown().
		SetLabel("Type: ").
		SetOptions([]string{"nc", "test1"}, func(t string, i int) {
			// TODO: From the data stuff that I'll import later, -
			// according to the index, get the correct payload and -
			// set the `listenerCommand`'s text.
		})

	listenerCopyButton := tview.NewButton("Copy").
		SetSelectedFunc(func() {
			// TODO: Probably do something like `xclip -selection clipboard`
			// out, _ := exec.Command("").Output()
			// fmt.Printf("%s", out)
		})

	listenerFlex := tview.NewFlex().
		AddItem(tview.NewTextView().SetText("Listener"), 0, 1, false).
		SetDirection(tview.FlexRow).
		AddItem(Spacer(), 0, 1, false).
		AddItem(listenerCommand, 0, 1, false).
		AddItem(tview.NewFlex().
			AddItem(listenerTypeDropdown, 0, 3, true).
			AddItem(listenerCopyButton, 0, 1, true), 0, 1, true)

	headerGrid.AddItem(ipPortFlex, 0, 0, 1, 1, 0, 0, true).
		AddItem(Spacer(), 0, 1, 1, 1, 0, 0, false).
		AddItem(listenerFlex, 0, 2, 1, 1, 0, 0, true)

	a.mainGrid.AddItem(title, 0, 0, 1, 1, 0, 0, false).
		AddItem(headerGrid, 1, 0, 1, 1, 0, 0, true)

	a.app.SetRoot(a.mainGrid, true)
}
