package main

import (
	"rvfs/ui"
)

func main() {
	app := ui.New()

	if err := app.Run(); err != nil {
		panic(err)
	}
}
