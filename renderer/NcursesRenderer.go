package renderer

import (
	gc "github.com/rthornton128/goncurses"
	"log"
	"gameoflife/model"
)

type NcursesRenderer struct {
	model.Renderer
	stdscr *gc.Window
	pad *gc.Pad
	err error
	width, height int
	universe *model.Universe
}

func NewNcursesRenderer(width int, height int) model.Renderer {
	renderer := &NcursesRenderer{}
	renderer.init(width, height)

	return renderer
}

func (renderer *NcursesRenderer) init(width int, height int) {
	stdscr, err := gc.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer gc.End()

	renderer.stdscr = stdscr

	renderer.width = width
	renderer.height = height

	gc.Cursor(0)
	gc.Echo(false)
	gc.HalfDelay(1)
}

func (renderer *NcursesRenderer) Draw(cell *model.Cell) {
	w, err := gc.NewWindow(1, 1, cell.GetY(), cell.GetX())
	if err != nil {
		log.Println("newBullet:", err)
	}

	if cell.IsAlive() {
		w.Print("0")
	}

	renderer.stdscr.Overlay(w);
}

func (renderer *NcursesRenderer) BeforeDraw() {
	renderer.stdscr.Erase()
}

func (renderer *NcursesRenderer) AfterDraw() {
	renderer.stdscr.Refresh()
}
