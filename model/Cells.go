package model

type Cells struct {
	width  int
	height int
	Cells  [][] *Cell
}

func NewCells(width int, height int) *Cells {
	cells := make([][] *Cell, width)

	for x := range cells {
		cells[x] = make([] *Cell, height)

		for y := range cells[x] {
			cells[x][y] = NewCell(x, y)
		}
	}

	newCells := &Cells{width, height, cells}

	for x := range cells {
		for y := range cells[x] {
			cells[x][y].Init(newCells)
		}
	}

	return newCells
}

func (cells *Cells) Get(x int, y int) *Cell {
	if x < 0 || y < 0 || x >= cells.width || y >= cells.height {
		return nil
	}

	return cells.Cells[x][y]
}

func (cells *Cells) Set(x int, y int, alive bool) *Cell {
	cell := cells.Cells[x][y]
	cell.alive = alive

	return cell
}
