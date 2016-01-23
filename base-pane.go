package main

import (
	"fmt"
	"strings"

	"github.com/grafov/gocui"
)

const minSplit = 2

var (
	splitY int
)

// base panel for exploring queries and log output
func basePane(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if splitY == 0 {
		splitY = maxY - 4
	}
	l, err := g.SetView("log", 0, 0, maxX-1, splitY)
	if err != nil {
		if err != gocui.ErrorUnkView {
			return err
		}
		l.Frame = false
		l.Autoscroll = true
		fmt.Fprintf(l, strings.Repeat("\n", splitY))
	}
	s, err := g.SetView("sql", 0, splitY, maxX-1, maxY-2)
	if err != nil {
		if err != gocui.ErrorUnkView {
			return err
		}
		s.Editable = true
		s.Wrap = true
		s.Frame = false
		fmt.Fprint(s, "")
	}
	m, err := g.SetView("menu", 0, maxY-2, maxX-1, maxY)
	if err != nil {
		if err != gocui.ErrorUnkView {
			return err
		}
		m.Frame = false
		fmt.Fprint(m, "F1 new       F2 edit      F5 copy to clpbrd      F8 delete        F10 exit")
	}
	g.SetCurrentView("sql")
	g.ShowCursor = true
	return nil
}
