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
	if v, err := g.SetView("hello", 0, 0, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "asc中文개를１９あるノン"
		v.Wrap = true
		fmt.Fprintln(v, "这是一段中文！！")
		fmt.Fprintln(v, "고개를 치켜들고 세상을 똑바로 바라보라")
		fmt.Fprintln(v, "村上春樹（１９４９－）である。小説、エッセイ、ノンフィクション")
		fmt.Fprintln(v, `这是一段中文1！！这是一段中文2！！这是一段中文3！！这是一段中文4！！这是一段中文5！！这是一段中文6！！这是一段中文7！！这是一段中文8！！这是一段中文9！！这是一段中文10！！这是一段中文11！！这是一段中文12！！`)
	}
	return nil
}
func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
