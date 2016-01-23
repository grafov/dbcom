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
		v.Frame = false
		v.Autoscroll = true
		for i := 0; i < maxY-1; i++ {
			fmt.Fprint(v, "\n")
		}
		fmt.Fprint(v, "stub > ")
	}
	m, err := g.SetView("menu", 0, maxY-2, maxX-1, maxY)
	if err != nil {
		if err != gocui.ErrorUnkView {
			return err
		}
		m.Frame = false
		fmt.Fprint(m, "F1 new       F2 edit      F5 copy to clpbrd      F8 delete        F10 exit")
	}
	return nil
}
