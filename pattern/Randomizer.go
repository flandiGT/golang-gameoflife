package pattern

import (
	"gameoflife/core"
	"math/rand"
	"time"
)

func CreateRandomUniverse(width int, height int) *core.Universe {
	universe := core.NewUniverse(width,height)

	rand.Seed(time.Now().UTC().UnixNano())

	for x:= 0; x < width; x++ {
		for y := 0; y < height; y++ {
			universe.Set(x, y, rand.Int() % 3 == 1)
		}
	}

	return universe
}
