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

// sql panel is for exploring queries and log output
func panelsLayout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	// Panels of sql explorer
	if splitY == 0 {
		splitY = maxY - 4
	}
	l, err := g.SetView("log", 0, 0, maxX-1, splitY)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		l.Frame = false
		l.Autoscroll = true
		l.BgColor = gocui.ColorBlack
		l.FgColor = gocui.ColorWhite
		fmt.Fprintf(l, strings.Repeat("\n", splitY))
	}
	s, err := g.SetView("sql", 0, splitY, maxX-1, maxY-2)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		s.Editable = true
		s.Wrap = true
		s.Frame = false
		s.BgColor = gocui.ColorBlack
		s.FgColor = gocui.ColorWhite
		fmt.Fprint(s, "")
	}

	// Two panels over sql explorer
	lp, err := g.SetView("lpanel", 0, 0, maxX/2-1, maxY-2)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		lp.Autoscroll = true
		lp.BgColor = gocui.ColorBlue
		lp.Highlight = true
	}
	rp, err := g.SetView("rpanel", maxX/2, 0, maxX-1, maxY-2)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		rp.Autoscroll = true
		rp.BgColor = gocui.ColorBlue
		rp.Highlight = true
	}

	mp, err := g.SetView("menu", 0, maxY-2, maxX-1, maxY)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		mp.Frame = false
		mp.BgColor = gocui.ColorBlack
		mp.FgColor = gocui.ColorYellow | gocui.AttrBold
		fmt.Fprint(mp, "F1 new       F2 edit      F5 copy      F8 delete        F10 exit")
	}
	return nil
}
