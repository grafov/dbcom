package main

import "github.com/grafov/gocui"

func initKeys(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlQ, gocui.ModNone, quit); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyF10, gocui.ModNone, quit); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyCtrlO, gocui.ModNone, switchPanels); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyF2, gocui.ModAlt, addSource); err != nil {
		return err
	}

	if err := g.SetKeybinding("lpanel", gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
		return err
	}
	if err := g.SetKeybinding("rpanel", gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
		return err
	}
	if err := g.SetKeybinding("lpanel", gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
		return err
	}
	if err := g.SetKeybinding("rpanel", gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
		return err
	}

	if err := g.SetKeybinding("sql", gocui.KeyArrowDown, gocui.ModNone, cursorDownInEditor); err != nil {
		return err
	}
	if err := g.SetKeybinding("sql", gocui.KeyEnter, gocui.ModAlt, execQuery); err != nil {
		return err
	}

	return nil
}
