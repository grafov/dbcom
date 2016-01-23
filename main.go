package main

import (
	"log"

	"github.com/grafov/gocui"
)

func main() {
	var err error

	// TODO show base pane
	g := gocui.NewGui()
	if err := g.Init(); err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetLayout(basePane)
	if err := initKeys(g); err != nil {
		log.Panicln(err)
	}

	// TODO show two panes over

	// TODO init db on demand

	err = g.MainLoop()
	if err != nil && err != gocui.Quit {
		log.Panicln(err)
	}
}
