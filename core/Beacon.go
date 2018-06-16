package core

func CreateBeacon() *Universe {
	universe := NewUniverse(6, 6)

	universe.Set(1, 1, true)
	universe.Set(1, 2, true)
	universe.Set(2, 1, true)
	universe.Set(2, 2, true)

	universe.Set(3, 3, true)
	universe.Set(3, 4, true)
	universe.Set(4, 3, true)
	universe.Set(4, 4, true)

	return universe
}
