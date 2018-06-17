package core

type Cell struct {
	x         int
	y         int
	Alive     bool
	aliveNext bool
}

func NewCell(x int, y int) *Cell {
	return &Cell{x, y, false, false}
}

func (cell *Cell) willLive(alive bool) {
	cell.aliveNext = alive
}

func (cell *Cell) next() {
	cell.Alive = cell.aliveNext
}