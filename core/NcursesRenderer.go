package core

import (
	gc "github.com/rthornton128/goncurses"
	"log"
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

func (renderer *NcursesRenderer) init(universe *Universe) {
	stdscr, err := gc.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer gc.End()

	renderer.stdscr = stdscr

	renderer.width = universe.width
	renderer.height = universe.height

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

func (renderer *NcursesRenderer) Render(universe *Universe) {
	if renderer.stdscr == nil {
		renderer.init(universe)
	}

	renderer.stdscr.Erase()

	for x := 0; x < universe.width; x++ {
		for y := 0; y < universe.height; y++ {

			renderer.drawCell(x, y, universe.cells[x][y]);
		}
	}

	renderer.stdscr.Refresh()
}
