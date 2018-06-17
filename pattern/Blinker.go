package pattern

import "gameoflife/core"

func CreateBlinker() *core.Universe {
	universe := core.NewUniverse(5, 5)

	universe.Set(2, 1, true)
	universe.Set(2, 2, true)
	universe.Set(2, 3, true)

	return universe
}
