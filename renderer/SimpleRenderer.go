package renderer

import (
	"fmt"
	"gameoflife/core"
)

type SimpleRenderer struct {
}

func NewSimpleRenderer() *SimpleRenderer {
	return &SimpleRenderer{}
}

func (renderer *SimpleRenderer) Render(universe *core.Universe) {
	for x := 0; x < universe.Width; x++ {
		for y := 0; y < universe.Height; y++ {
			if universe.Cells[x][y].Alive {
				fmt.Print("0")
			} else {
				fmt.Print(".")
			}
		}

		fmt.Print("\n")
	}

	fmt.Print("-------------------------\n")
}