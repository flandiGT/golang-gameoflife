package pattern

import "gameoflife/core"

func CreatePentadecathlon() *core.Universe {
	universe := core.NewUniverse(11, 18)

	universe.Set(4, 3, true)
	universe.Set(5, 3, true)
	universe.Set(6, 3, true)

	universe.Set(5, 4, true)
	universe.Set(5, 5, true)

	universe.Set(4, 6, true)
	universe.Set(5, 6, true)
	universe.Set(6, 6, true)


	universe.Set(4, 8, true)
	universe.Set(5, 8, true)
	universe.Set(6, 8, true)
	universe.Set(4, 9, true)
	universe.Set(5, 9, true)
	universe.Set(6, 9, true)


	universe.Set(4, 11, true)
	universe.Set(5, 11, true)
	universe.Set(6, 11, true)

	universe.Set(5, 12, true)
	universe.Set(5, 13, true)

	universe.Set(4, 14, true)
	universe.Set(5, 14, true)
	universe.Set(6, 14, true)

	return universe
}
