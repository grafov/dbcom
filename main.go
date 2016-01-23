package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gophergala2016/dbcom/db"
	"github.com/grafov/gocui"
	_ "github.com/jackc/pgx"
)

func main() {
	var err error

	// TODO show base pane
	g := gocui.NewGui()
	if err = g.Init(); err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetLayout(twoPanels)

	if err = initKeys(g); err != nil {
		log.Panicln(err)
	}

	// TODO show two panes over

	// TODO init db on demand

	// for stub just init test db with hardcoded values
	if err = db.Add("test", "mysql", "root:root@/test"); err != nil {
		log.Panicln(err)
	}

	if err = g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)

	}
}
