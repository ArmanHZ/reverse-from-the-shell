package main

import (
	"flag"
	"rvfs/ui"
)

func main() {
	var isColorFix bool
	flag.BoolVar(&isColorFix, "color-fix", false, "Set this to true if you are having color issues.")

	flag.Parse()

	app := ui.New(isColorFix)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
