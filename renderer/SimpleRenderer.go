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
			cell := universe.Get(x, y)
			if cell.IsAlive() {
				fmt.Print("0")
			} else {
				fmt.Print(".")
			}
		}

		fmt.Print("\n")
	}

	fmt.Print("-------------------------\n")
}