package main

import "github.com/grafov/gocui"

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.Quit
}

func refresh(g *gocui.Gui, v *gocui.View) error {
	return g.Flush()
}

// cursorDown enlarge sql view when a new line added
func cursorDown(g *gocui.Gui, v *gocui.View) error {
	x, y := v.Cursor()
	_, sizeY := v.Size()
	if y >= sizeY-1 {
		if splitY == minSplit {
			return nil
		}
		splitY--
		l, _ := g.View("log")
		l.Resize(0, -1)
		v.Resize(0, 1)
		v.SetOrigin(0, 0)
		v.SetCursor(x, y+1)
	}
	return nil
}
