package main

import (
	// "fmt"
	// "os/exec"

	"rvfs/ui"
)

func main() {
	app := ui.New()

	if err := app.Run(); err != nil {
		panic(err)
	}
}
