package core

import (
	"container/list"
)

type Universe struct {
	Width  int
	Height int
	cells  *Cells
	renderer Renderer
}

func NewUniverse(width int, height int) *Universe {
	return &Universe{width,height,NewCells(width, height), nil}
}

func (universe *Universe) Set(x int, y int, alive bool) {
	universe.cells.Set(x, y, alive)
}

func (universe *Universe) Get(x int, y int) *Cell {
	return universe.cells.Get(x, y)
}

func (Universe *Universe) SetRenderer(renderer Renderer) {
	Universe.renderer = renderer
}

func (universe *Universe) Next()  {
	cellsToBeChanged := list.New()

	for x := 0; x < universe.Width; x++ {
		for y := 0; y < universe.Height; y++ {
			cell := universe.cells.Get(x, y)

			if cell.estimateChange() {
				cellsToBeChanged.PushBack(cell)
			}
		}
	}

	universe.renderer.BeforeDraw()

	element := cellsToBeChanged.Front()
	for ; element != nil; element = element.Next() {
		cell := element.Value.(*Cell)

		cell.next()
		universe.renderer.Draw(cell)
	}

	universe.renderer.AfterDraw()
}
