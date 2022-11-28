package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jroimartin/gocui"
	"github.com/sirupsen/logrus"
)

func main() {
	logfile, e := os.OpenFile("ch.log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if e != nil {
		fmt.Println("failed open log file:", e)
		return
	}
	defer logfile.Close()
	logrus.SetOutput(logfile)
	logrus.SetLevel(logrus.TraceLevel)
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("hello", 0, 0, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "这是一段中文！！") // 8
	}
	return nil
}
func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
