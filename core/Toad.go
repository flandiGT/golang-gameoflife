package core

func CreateToad() *Universe {
	universe := NewUniverse(6, 6)

	universe.Set(3, 1, true)
	universe.Set(4, 2, true)
	universe.Set(4, 3, true)

	universe.Set(1, 2, true)
	universe.Set(1, 3, true)
	universe.Set(2, 4, true)

	return universe
}
