package core

func CreateGosperGliderGun() *Universe {
	universe := NewUniverse(38, 11)

	universe.Set(1, 5, true)
	universe.Set(1, 6, true)
	universe.Set(2, 5, true)
	universe.Set(2, 6, true)


	universe.Set(11, 5, true)
	universe.Set(11, 6, true)
	universe.Set(11, 7, true)

	universe.Set(12, 4, true)
	universe.Set(12, 8, true)

	universe.Set(13, 3, true)
	universe.Set(13, 9, true)

	universe.Set(14, 3, true)
	universe.Set(14, 9, true)

	universe.Set(15, 6, true)

	universe.Set(16, 4, true)
	universe.Set(16, 8, true)

	universe.Set(17, 5, true)
	universe.Set(17, 6, true)
	universe.Set(17, 7, true)

	universe.Set(18, 6, true)


	universe.Set(21, 3, true)
	universe.Set(21, 4, true)
	universe.Set(21, 5, true)

	universe.Set(22, 3, true)
	universe.Set(22, 4, true)
	universe.Set(22, 5, true)

	universe.Set(23, 2, true)
	universe.Set(23, 6, true)

	universe.Set(25, 1, true)
	universe.Set(25, 2, true)
	universe.Set(25, 6, true)
	universe.Set(25, 7, true)



	universe.Set(35, 3, true)
	universe.Set(35, 4, true)
	universe.Set(36, 3, true)
	universe.Set(36, 4, true)

	return universe
}
