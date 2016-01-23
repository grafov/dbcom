package main

import "github.com/grafov/gocui"

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.Quit
}
