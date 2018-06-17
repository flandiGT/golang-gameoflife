package core

import (
	"container/list"
)

type Cell struct {
	x             int
	y             int
	alive         bool
	willLive      bool
	neighborCount int
	neighbors     [] *Cell
}

func NewCell(x int, y int) *Cell {
	neighbors := make([] *Cell, 8)
	return &Cell{
		x,
		y,
		false,
		false,
		0,
		neighbors}
}

func (cell *Cell) addNeighbor(neighbor *Cell) {
	if neighbor != nil {
		cell.neighbors[cell.neighborCount] = neighbor
		cell.neighborCount++
	}
}

func (cell *Cell) Init(cells *Cells) {
	// North-West
	cell.addNeighbor(cells.Get(cell.x - 1, cell.y - 1))

	// North
	cell.addNeighbor(cells.Get(cell.x - 1, cell.y))


	// North-East
	cell.addNeighbor(cells.Get(cell.x - 1, cell.y + 1))


	// West
	cell.addNeighbor(cells.Get(cell.x, cell.y - 1))

	// East
	cell.addNeighbor(cells.Get(cell.x, cell.y + 1))


	// South-West
	cell.addNeighbor(cells.Get(cell.x + 1, cell.y - 1))


	// South
	cell.addNeighbor(cells.Get(cell.x + 1, cell.y))


	// South-East
	cell.addNeighbor(cells.Get(cell.x + 1, cell.y + 1))
}

func (cell *Cell) IsAlive() bool {
	return cell.alive
}

func (cell *Cell) GetX() int {
	return cell.x
}

func (cell *Cell) GetY() int {
	return cell.y
}

func (cell *Cell) estimateBornNeighbors(changedCells *list.List) {
	for i := 0; i < cell.neighborCount; i++ {
		neighbor := cell.neighbors[i]

		if neighbor.IsAlive() == false && neighbor.willLive == false {

			willBeBorn := neighbor.estimateBirth()

			if willBeBorn {
				neighbor.willLive = true
				changedCells.PushBack(neighbor)
			}
		}
	}
}

func (cell *Cell) countLivingNeighborsUntil(limit int) int {
	livingNeighbors := 0

	for i := 0; i < cell.neighborCount && livingNeighbors < limit; i++ {
		neighbor := cell.neighbors[i]

		if neighbor.IsAlive() {
			livingNeighbors++
		}
	}

	return livingNeighbors
}

func (cell *Cell) estimateBirth() bool {
	livingNeighbors := cell.countLivingNeighborsUntil(9)

	// Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
	if livingNeighbors == 3 {
		return true
	}

	return false
}

func (cell *Cell) stayingAlive() bool {
	livingNeighbors := cell.countLivingNeighborsUntil(9)

	// 	Any live cell with fewer than two live neighbors dies, as if by under population.
	if livingNeighbors < 2 {
		return false
	}

	// Any live cell with two or three live neighbors lives on to the next generation.
	if livingNeighbors == 2 || livingNeighbors == 3 {
		return true
	}

	// Any live cell with more than three live neighbors dies, as if by overpopulation.
	if livingNeighbors > 3 {
		return false
	}

	// should never happen -> exception?
	return false
}
