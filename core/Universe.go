package core

type Universe struct {
	Width               int
	Height              int
	Cells               [][]bool
	nextGenerationCells [][]bool
}

func NewUniverse(width int, height int) *Universe {
	var cells = make([][]bool, width)
	var nextGenerationCells = make([][]bool, width)

	for x := range cells {
		cells[x] = make([]bool, height)
		nextGenerationCells[x] = make([]bool, height)

		for y := range cells[x] {
			cells[x][y] = false
			nextGenerationCells[x][y] = false
		}
	}

	return &Universe{width,height,cells, nextGenerationCells}
}

func (universe *Universe) Get(x int, y int) bool {
	if x < 0 || y < 0 || x >= universe.Width || y >= universe.Height {
		return false
	}

	return universe.Cells[x][y]
}

func (universe *Universe) Set(x int, y int, isLiving bool) {
	universe.Cells[x][y] = isLiving
}

func (universe *Universe) Next()  {

	for x := 0; x < universe.Width; x++ {
		for y := 0; y < universe.Height; y++ {
			livingNeighbors := universe.countLivingNeighbors(x, y)
			isLiving := universe.Cells[x][y]

			universe.nextGenerationCells[x][y] = universe.willLive(isLiving, livingNeighbors)
		}
	}

	temporaryCells := universe.Cells
	universe.Cells = universe.nextGenerationCells
	universe.nextGenerationCells = temporaryCells
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