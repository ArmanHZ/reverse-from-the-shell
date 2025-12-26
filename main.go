package main

import (
	// "fmt"
	// "os/exec"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// TODO: Will be useful later on defining the vim-style movement keys.
func initControlls(app *tview.Application, row int, col int, cells [][]tview.Primitive) {
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlH:
			col = max(0, col-1)
		case tcell.KeyCtrlL:
			col = min(len(cells[0])-1, col+1)
		case tcell.KeyCtrlK:
			row = max(0, row-1)
		case tcell.KeyCtrlJ:
			row = min(len(cells)-1, row+1)
		default:
			return event
		}

		app.SetFocus(cells[row][col])
		return nil
	})
}

func spacer() *tview.Box {
	return tview.NewBox()
}


func main() {
	app := tview.NewApplication()

	mainGrid := tview.NewGrid().
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
		AddItem(spacer(), 0, 1, false).
		AddItem(tview.NewFlex().
			AddItem(ipField, 0, 4, true).
			AddItem(spacer(), 0, 1, false).
			AddItem(portField, 0, 4, true), 0, 1, true)
	
	listenerCommand := tview.NewTextView().
		SetText("nc -lvnp " + portField.GetText()).
		SetTextAlign(tview.AlignCenter)
	
	portField.SetChangedFunc(func(text string) {
		listenerCommand.SetText("nc -lvnp " + text)
	})

	listenerTypeDropdown := tview.NewDropDown().
		SetLabel("Type: ").
		SetOptions([]string{"nc", "test1"}, func(t string, i int){
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
		AddItem(spacer(), 0, 1, false).
		AddItem(listenerCommand, 0, 1, false).
		AddItem(tview.NewFlex().
			AddItem(listenerTypeDropdown, 0, 3, true).
			AddItem(listenerCopyButton, 0, 1, true), 0, 1, true)
	

	headerGrid.AddItem(ipPortFlex, 0, 0, 1, 1, 0, 0, true).
		AddItem(spacer(), 0, 1, 1, 1, 0, 0, false).
		AddItem(listenerFlex, 0, 2, 1, 1, 0, 0, true)
	
	mainGrid.AddItem(title, 0, 0, 1, 1, 0, 0, false).
		AddItem(headerGrid, 1, 0, 1, 1, 0, 0, true)
	
	
	// row, col := 0, 0
	// cells := [][]tview.Primitive{}

	// initControlls(app, row, col, cells)

	// Line with SetFocus for testing purposes.
	// err := app.SetRoot(mainGrid, true).SetFocus(testButton).Run(); if err != nil {
	err := app.SetRoot(mainGrid, true).Run(); if err != nil {
		panic(err)
	}
}
