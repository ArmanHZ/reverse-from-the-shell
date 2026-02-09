package main

import (
	"flag"
	"rvfs/ui"
)

func main() {
	var isColorFix bool
	flag.BoolVar(&isColorFix, "color-fix", false, "Set this to true if you are having color issues.")

	var userIp string
	flag.StringVar(&userIp, "i", "", "Set this if you don't want to set it from the TUI.")

	flag.Parse()

	app := ui.New(isColorFix, userIp)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
