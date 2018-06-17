package renderer

import (
	gc "github.com/rthornton128/goncurses"
	"log"
	"gameoflife/core"
)

type NcursesRenderer struct {
	stdscr *gc.Window
	pad *gc.Pad
	err error
	width, height int
}

func NewNcursesRenderer() *NcursesRenderer {
	return &NcursesRenderer{}
}

func (renderer *NcursesRenderer) init(universe *core.Universe) {
	stdscr, err := gc.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer gc.End()

	renderer.stdscr = stdscr

	renderer.width = universe.Width
	renderer.height = universe.Height

	gc.Cursor(0)
	gc.Echo(false)
	gc.HalfDelay(1)
}

func (renderer *NcursesRenderer) drawCell(x int, y int, alive bool) {
	w, err := gc.NewWindow(1, 1, y, x)
	if err != nil {
		log.Println("newBullet:", err)
	}

	if alive {
		w.Print("0")
	} else {
		w.Print(" ")
	}

	renderer.stdscr.Overlay(w);
}

func (renderer *NcursesRenderer) Render(universe *core.Universe) {
	if renderer.stdscr == nil {
		renderer.init(universe)
	}

	renderer.stdscr.Erase()

	for x := 0; x < universe.Width; x++ {
		for y := 0; y < universe.Height; y++ {
			cell := universe.Get(x, y)
			renderer.drawCell(x, y, cell.IsAlive());
		}
	}

	renderer.stdscr.Refresh()
}
