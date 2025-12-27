package ui

func (a *App) bindEvents() {
	a.portField.SetChangedFunc(func(text string) {
		a.listenerCommand.SetText("nc -lvnp " + text)
	})

	// TODO: From the data stuff that I'll import later, -
	// according to the index, get the correct payload and -
	// set the `listenerCommand`'s text.

	// TODO: Probably do something like `xclip -selection clipboard`
	// out, _ := exec.Command("").Output()
	// fmt.Printf("%s", out)
}
