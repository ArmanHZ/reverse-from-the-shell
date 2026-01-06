package ui

import (
	// "github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (a *App) buildHeader() tview.Primitive {
	headerGrid := tview.NewGrid().
		SetRows(0).
		SetColumns(0, 4, 0)

	a.ipField = tview.NewInputField().
		SetLabel("IP: ").
		SetText("10.10.10.10")

	a.portField = tview.NewInputField().
		SetLabel("Port: ").
		SetText("9001")

	ipPortFlex := tview.NewFlex().
		AddItem(tview.NewTextView().SetText("IP & Port"), 1, 0, false).
		SetDirection(tview.FlexRow).
		AddItem(Spacer(), 1, 0, false).
		AddItem(tview.NewFlex().
			AddItem(a.ipField, 0, 4, true).
			AddItem(Spacer(), 0, 1, false).
			AddItem(a.portField, 0, 4, true), 0, 1, true)

	a.listenerCommand = tview.NewTextView().
		SetText("nc -lvnp " + a.portField.GetText()).
		SetTextAlign(tview.AlignCenter)

	a.listenerTypeSelect = tview.NewDropDown().
		SetLabel("Type: ")

	// TODO: Move to events when fully implementing.
	listenerCopyButton := tview.NewButton("Copy").
		SetSelectedFunc(func() {
			// TODO: Probably do something like `xclip -selection clipboard`
			// out, _ := exec.Command("").Output()
			// fmt.Printf("%s", out)
		})

	listenerFlex := tview.NewFlex().
		AddItem(tview.NewTextView().SetText("Listener"), 1, 0, false).
		SetDirection(tview.FlexRow).
		AddItem(Spacer(), 1, 0, false).
		AddItem(a.listenerCommand, 0, 1, false).
		AddItem(tview.NewFlex().
			AddItem(a.listenerTypeSelect, 0, 3, true).
			AddItem(listenerCopyButton, 0, 1, true), 1, 0, true)

	headerGrid.AddItem(ipPortFlex, 0, 0, 1, 1, 0, 0, true).
		AddItem(Spacer(), 0, 1, 1, 1, 0, 0, false).
		AddItem(listenerFlex, 0, 2, 1, 1, 0, 0, true)

	return headerGrid
}

func (a *App) buildTabs() *tview.Flex {
	reverseTab := tview.NewButton("Reverse")
	bindTab := tview.NewButton("Bind")
	msfvenomTab := tview.NewButton("MSFVenom")
	hoaxTab := tview.NewButton("HoaxShell")

	a.tabButtons = append(a.tabButtons, reverseTab, bindTab, msfvenomTab, hoaxTab)

	tabsFlex := tview.NewFlex().
		AddItem(reverseTab, 0, 1, true).
		AddItem(Spacer(), 1, 0, false).
		AddItem(bindTab, 0, 1, true).
		AddItem(Spacer(), 1, 0, false).
		AddItem(msfvenomTab, 0, 1, true).
		AddItem(Spacer(), 1, 0, false).
		AddItem(hoaxTab, 0, 1, true).
		AddItem(Spacer(), 0, 3, false)

	return tabsFlex
}

func (a *App) buildMainContent() *tview.Flex {
	mainContentFlex := tview.NewFlex().SetDirection(tview.FlexRow)
	tabs := a.buildTabs()

	// T00d00: Refac these l8r
	targetOsSelect := tview.NewDropDown().SetLabel("OS: ").
		SetOptions([]string{"First", "Second", "Third", "Fourth", "Fifth"}, nil) // Dummdumm data for testing.
	payloadSearchField := tview.NewInputField().SetLabel("Name: ")

	mainContentControls := tview.NewFlex().
		AddItem(targetOsSelect, 0, 1, true).
		AddItem(Spacer(), 1, 0, false).
		AddItem(payloadSearchField, 0, 4, true).
		AddItem(Spacer(), 0, 5, false)

	a.reverseShellSelect = tview.NewTable().
		SetBorders(true).
		SetSelectable(true, false)

	// Dummy data
	items := []string{
		"First item",
		"Second item",
		"Third item",
		"Fourth item",
		"Fifth item",
	}

	for row, text := range items {
		cell := tview.NewTableCell(text).
			SetAlign(tview.AlignLeft)
			// SetExpansion(1)

		a.reverseShellSelect.SetCell(row, 0, cell)
	}

	mainContentData := tview.NewFlex().
		AddItem(a.reverseShellSelect, 0, 1, true)

	mainContentFlex.AddItem(tabs, 1, 0, true).
		AddItem(Spacer(), 1, 0, false).
		AddItem(mainContentControls, 1, 0, true).
		AddItem(Spacer(), 1, 0, false).
		AddItem(mainContentData, 0, 1, true)

	return mainContentFlex
}

func (a *App) buildUI() {
	a.mainGrid = tview.NewGrid().
		SetRows(1, 8, 0). // Value 8 for the second row seems to be a good value to fit the longest listener string.
		SetColumns(0).
		SetBorders(true)

	title := tview.NewTextView().
		SetDynamicColors(true).
		SetText("Reverse Shell Generator").
		SetTextAlign(tview.AlignCenter)

	header := a.buildHeader()
	// tabs := a.buildTabs()
	mainContent := a.buildMainContent()

	a.mainGrid.AddItem(title, 0, 0, 1, 1, 0, 0, false).
		AddItem(header, 1, 0, 1, 1, 0, 0, true).
		AddItem(mainContent, 2, 0, 1, 1, 0, 0, true)
		// AddItem(tabs, 2, 0, 1, 1, 0, 0, true).
		// AddItem(Spacer(), 3, 0, 1, 1, 0, 0, false)

	a.app.SetRoot(a.mainGrid, true)
}
