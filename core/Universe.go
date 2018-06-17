package core

type Universe struct {
	Width               int
	Height              int
	Cells               [][] *Cell
}

func NewUniverse(width int, height int) *Universe {
	var cells = make([][] *Cell, width)

	for x := range cells {
		cells[x] = make([] *Cell, height)

		for y := range cells[x] {
			cells[x][y] = NewCell(x, y)
		}
	}

	return &Universe{width,height,cells}
}

func (universe *Universe) Get(x int, y int) bool {
	if x < 0 || y < 0 || x >= universe.Width || y >= universe.Height {
		return false
	}

	return universe.Cells[x][y].Alive
}

func (universe *Universe) Set(x int, y int, isLiving bool) {
	universe.Cells[x][y].Alive = isLiving
}

func (universe *Universe) Next()  {

	for x := 0; x < universe.Width; x++ {
		for y := 0; y < universe.Height; y++ {
			livingNeighbors := universe.countLivingNeighbors(x, y)
			cell := universe.Cells[x][y]

			willLive := universe.willLive(cell.Alive, livingNeighbors)
			cell.willLive(willLive)
		}
	}

	for x := 0; x < universe.Width; x++ {
		for y := 0; y < universe.Height; y++ {
			cell := universe.Cells[x][y]
			cell.next()
		}
	}
}

func (universe *Universe) countLivingNeighbors(x int, y int) int {
	livingNeighbors := 0

	// North-West
	if universe.Get(x - 1, y - 1) {
		livingNeighbors++
	}

	// North
	if universe.Get(x - 1, y) {
		livingNeighbors++
	}

	// North-East
	if universe.Get(x - 1, y + 1) {
		livingNeighbors++
	}

	// West
	if universe.Get(x, y - 1) {
		livingNeighbors++
	}

	// East
	if universe.Get(x, y + 1) {
		livingNeighbors++
	}

	// South-West
	if universe.Get(x + 1, y - 1) {
		livingNeighbors++
	}

	// South
	if universe.Get(x + 1, y) {
		livingNeighbors++
	}

	// South-East
	if universe.Get(x + 1, y + 1) {
		livingNeighbors++
	}

	return livingNeighbors
}

func (universe *Universe) willLive(isLiving bool, livingNeighbors int) bool {

	if isLiving {
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
	} else {
		// Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
		if livingNeighbors == 3 {
			return true
		}
	}

	return isLiving
}