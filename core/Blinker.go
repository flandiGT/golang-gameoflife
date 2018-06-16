package core

func CreateBlinker() *Universe {
	universe := NewUniverse(5, 5)

	universe.Set(2, 1, true)
	universe.Set(2, 2, true)
	universe.Set(2, 3, true)

	return universe
}
