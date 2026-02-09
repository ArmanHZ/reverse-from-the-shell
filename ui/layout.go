package ui

import (
	"rvfs/data"
	"unicode"

	"github.com/rivo/tview"
)

func (a *App) buildHeader() tview.Primitive {
	headerGrid := tview.NewGrid().
		SetRows(0).
		SetColumns(0, 4, 0)

	// TODO: Later have the option to get this from cmdline and use placeholder.
	a.ipField = tview.NewInputField().
		SetLabel("IP: ").
		SetText("10.10.10.10").
		// Temporary solution for letters. Will use regex later.
		SetAcceptanceFunc(func(textToCheck string, lastChar rune) bool {
			if unicode.IsDigit(lastChar) || unicode.IsPunct(lastChar) {
				return true
			}

			return false
		})
		// SetPlaceholder("10.10.10.10")
	a.registerFocusable(a.ipField)

	a.portField = tview.NewInputField().
		SetLabel("Port: ").
		SetText("9001").
		SetAcceptanceFunc(tview.InputFieldInteger)
	a.registerFocusable(a.portField)

	ipPortFlex := tview.NewFlex().
		AddItem(tview.NewTextView().SetText("IP & Port"), 1, 0, false).
		SetDirection(tview.FlexRow).
		AddItem(Spacer(), 1, 0, false).
		AddItem(tview.NewFlex().
			AddItem(a.ipField, 0, 4, true).
			AddItem(Spacer(), 0, 1, false).
			AddItem(a.portField, 0, 4, true), 0, 1, true)

	a.listenerCommand = tview.NewTextView().
		SetTextAlign(tview.AlignCenter)

	a.listenerTypeSelect = tview.NewDropDown().
		SetLabel("Type: ")
	a.registerFocusable(a.listenerTypeSelect)

	a.listenerCopyButton = tview.NewButton("Copy")
	a.registerFocusable(a.listenerCopyButton) // TODO: Need to make the button global as well.

	listenerFlex := tview.NewFlex().
		AddItem(tview.NewTextView().SetText("Listener"), 1, 0, false).
		SetDirection(tview.FlexRow).
		AddItem(Spacer(), 1, 0, false).
		AddItem(a.listenerCommand, 0, 1, false).
		AddItem(tview.NewFlex().
			AddItem(a.listenerTypeSelect, 0, 3, true).
			AddItem(a.listenerCopyButton, 0, 1, true), 1, 0, true)

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

	for _, v := range a.tabButtons {
		a.registerFocusable(v)
	}

	// XXX: Disabling these for now. Will implement l8r.
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
	// tabs := a.buildTabs()

	a.targetOsTypeSelect = tview.NewDropDown().SetLabel("OS: ")
	a.registerFocusable(a.targetOsTypeSelect)

	payloadSearchField := tview.NewInputField().SetLabel("Name: ")
	payloadSearchField.SetPlaceholder("Will be implemented later...")
	// TODO: Will implement later.
	// a.registerFocusable(payloadSearchField)

	mainContentControls := tview.NewFlex().
		AddItem(a.targetOsTypeSelect, 0, 1, true).
		AddItem(Spacer(), 1, 0, false).
		AddItem(payloadSearchField, 0, 4, true).
		AddItem(Spacer(), 0, 5, false)

	a.shellCommandTable = tview.NewTable().
		SetBorders(true).
		SetSelectable(true, false)
	a.registerFocusable(a.shellCommandTable)

	a.shellPayloadDisplay = tview.NewTextView().
		SetDynamicColors(true)

	mainContentData := tview.NewGrid().
		SetRows(0).
		SetColumns(25, 0)

	// TODO: This section needs refactoring.
	a.shellTypeSelect = tview.NewDropDown().
		SetLabel("Shell: ").
		SetOptions(data.ShellTypes, nil).
		SetCurrentOption(2) // Bash by default
	a.registerFocusable(a.shellTypeSelect)

	encodingSelect := tview.NewDropDown().
		SetLabel("Encoding: ").
		SetOptions([]string{"Will implement l8r"}, nil).
		SetCurrentOption(0)

	a.shellPayloadCopyButton = tview.NewButton("Copy")
	a.registerFocusable(a.shellPayloadCopyButton)

	shellPayloadDisplayOptions := tview.NewFlex().
		AddItem(a.shellTypeSelect, 0, 1, true).
		AddItem(encodingSelect, 0, 1, true).
		AddItem(a.shellPayloadCopyButton, 10, 0, true)

	shellPayloadDisplayGrid := tview.NewGrid().
		SetRows(0, 1).
		SetColumns(0).
		AddItem(a.shellPayloadDisplay, 0, 0, 1, 1, 0, 0, true).
		AddItem(shellPayloadDisplayOptions, 1, 0, 1, 1, 0, 0, true)

	mainContentData.AddItem(a.shellCommandTable, 0, 0, 1, 1, 0, 0, true).
		AddItem(shellPayloadDisplayGrid, 0, 1, 1, 1, 0, 0, true)

	// FIXME: Temporarily removed the tabs
	mainContentFlex. // AddItem(tabs, 1, 0, true).
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
	mainContent := a.buildMainContent()

	a.mainGrid.AddItem(title, 0, 0, 1, 1, 0, 0, false).
		AddItem(header, 1, 0, 1, 1, 0, 0, true).
		AddItem(mainContent, 2, 0, 1, 1, 0, 0, true)

	a.app.SetRoot(a.mainGrid, true)
}
