package core

import (
	"fmt"
)

type SimpleRenderer struct {
}

func NewSimpleRenderer() *SimpleRenderer {
	return &SimpleRenderer{}
}

func (renderer *SimpleRenderer) Render(universe *Universe) {
	for x := 0; x < universe.width; x++ {
		for y := 0; y < universe.height; y++ {
			if universe.cells[x][y] {
				fmt.Print("0")
			} else {
				fmt.Print(".")
			}
		}

		fmt.Print("\n")
	}

	fmt.Print("-------------------------\n")
}