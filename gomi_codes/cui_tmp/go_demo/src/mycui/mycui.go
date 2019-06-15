package main

import (
	"fmt"
	"log"
	"github.com/jroimartin/gocui"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	// if err := g.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
	// 	log.Panicln(err)
	// }

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	if v, err := g.SetView("view 1", 2, 2, 22, 7); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		
		items := map[int]string{0:"Item 1" , 1:"Item 2"}
		for key, value := range items {
			fmt.Fprintf(v, "%v:%v\n", key, value)
		}
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

// func cursorDown(g *gocui.Gui, v *gocui.View) error {
// 	fmt.Println("cursorDown")
// 	if v != nil {
// 		cx, cy := v.Cursor()
// 		if err := v.SetCursor(cx, cy+1); err != nil {
// 			ox, oy := v.Origin()
// 			if err := v.SetOrigin(ox, oy+1); err != nil {
// 				return err
// 			}
// 		}
// 	}
// 	return nil
// }