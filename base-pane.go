package main

import (
	"fmt"

	"github.com/grafov/gocui"
)

// base panel for exploring queries and log output
func basePane(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	v, err := g.SetView("root", 0, 0, maxX-1, maxY-1)
	if err != nil {
		if err != gocui.ErrorUnkView {
			return err
		}
		for i := 0; i < maxY-2; i++ {
			fmt.Fprintln(v, "")
		}
		fmt.Fprintln(v, "stub > ")
	}
	return nil
}
