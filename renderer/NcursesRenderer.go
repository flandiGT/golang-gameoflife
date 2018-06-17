package renderer

import (
	gc "github.com/rthornton128/goncurses"
	"log"
	"gameoflife/core"
)

type NcursesRenderer struct {
	core.Renderer
	stdscr *gc.Window
	pad *gc.Pad
	err error
	width, height int
	universe *core.Universe
}

func NewNcursesRenderer(width int, height int) *NcursesRenderer {
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

func (renderer *NcursesRenderer) Draw(cell *core.Cell) {
	w, err := gc.NewWindow(1, 1, cell.GetY(), cell.GetX())
	if err != nil {
		log.Println("newBullet:", err)
	}

	if cell.IsAlive() {
		w.Print("0")
	} else {
		w.Print(" ")
	}

	renderer.stdscr.Overlay(w);
}

func (renderer *NcursesRenderer) BeforeDraw() {
	renderer.stdscr.Erase()
}

func (renderer *NcursesRenderer) AfterDraw() {
	renderer.stdscr.Refresh()
}

//func (renderer *NcursesRenderer) Render() {
//	renderer.stdscr.Erase()
//
//	for x := 0; x < renderer.universe.Width; x++ {
//		for y := 0; y < renderer.universe.Height; y++ {
//			cell := renderer.universe.Get(x, y)
//			renderer.Draw(cell);
//		}
//	}
//
//	renderer.stdscr.Refresh()
//}
