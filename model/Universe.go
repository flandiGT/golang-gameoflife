package model

import (
	"container/list"
)

type Universe struct {
	Width  int
	Height int
	cells  *Cells
	livingCells *list.List
	renderer Renderer
}

func NewUniverse(width int, height int) *Universe {
	return &Universe{width,
	height,
	NewCells(width, height),
	&list.List{},
	nil}
}

func (universe *Universe) Set(x int, y int, alive bool) {
	cell := universe.cells.Set(x, y, alive)

	if alive {
		universe.livingCells.PushBack(cell)
	}
}

func (universe *Universe) Get(x int, y int) *Cell {
	return universe.cells.Get(x, y)
}

func (Universe *Universe) SetRenderer(renderer Renderer) {
	Universe.renderer = renderer
}

func (universe *Universe) Next()  {
	nextLivingCells := list.New()
	dieingCells := list.New()

	universe.renderer.BeforeDraw()

	for element := universe.livingCells.Front(); element != nil; element = element.Next() {
		cell := element.Value.(*Cell)
		universe.renderer.Draw(cell)

		if cell.stayingAlive() {
			nextLivingCells.PushBack(cell)
		} else {
			dieingCells.PushBack(cell)
		}

		cell.estimateBornNeighbors(nextLivingCells)
	}

	universe.renderer.AfterDraw()
	universe.livingCells = nextLivingCells

	for element := universe.livingCells.Front(); element != nil; element = element.Next() {
		cell := element.Value.(*Cell)
		cell.alive = true
		cell.willLive = false
	}

	for element := dieingCells.Front(); element != nil; element = element.Next() {
		cell := element.Value.(*Cell)
		cell.alive = false
		cell.willLive = false
	}
}
