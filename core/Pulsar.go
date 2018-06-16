package core

func CreatePulsar() *Universe {
	universe := NewUniverse(17, 17)

	universe.Set(5, 1, true)
	universe.Set(5, 2, true)
	universe.Set(5, 3, true)
	universe.Set(6, 3, true)

	universe.Set(11, 1, true)
	universe.Set(11, 2, true)
	universe.Set(11, 3, true)
	universe.Set(10, 3, true)

	universe.Set(1, 5, true)
	universe.Set(2, 5, true)
	universe.Set(3, 5, true)
	universe.Set(3, 6, true)

	universe.Set(13, 5, true)
	universe.Set(14, 5, true)
	universe.Set(15, 5, true)
	universe.Set(13, 6, true)

	universe.Set(1, 11, true)
	universe.Set(2, 11, true)
	universe.Set(3, 11, true)
	universe.Set(3, 10, true)

	universe.Set(13, 11, true)
	universe.Set(14, 11, true)
	universe.Set(15, 11, true)
	universe.Set(13, 10, true)

	universe.Set(5, 13, true)
	universe.Set(5, 14, true)
	universe.Set(5, 15, true)
	universe.Set(6, 13, true)

	universe.Set(11, 13, true)
	universe.Set(11, 14, true)
	universe.Set(11, 15, true)
	universe.Set(10, 13, true)



	universe.Set(6, 5, true)
	universe.Set(7, 5, true)
	universe.Set(7, 6, true)

	universe.Set(5, 6, true)
	universe.Set(5, 7, true)
	universe.Set(6, 7, true)


	universe.Set(9, 5, true)
	universe.Set(9, 6, true)
	universe.Set(10, 5, true)

	universe.Set(11, 6, true)
	universe.Set(11, 7, true)
	universe.Set(10, 7, true)


	universe.Set(5, 9, true)
	universe.Set(5, 10, true)
	universe.Set(6, 9, true)

	universe.Set(7, 10, true)
	universe.Set(7, 11, true)
	universe.Set(6, 11, true)


	universe.Set(9, 10, true)
	universe.Set(10, 10, true)
	universe.Set(10, 11, true)

	universe.Set(10, 9, true)
	universe.Set(11, 9, true)
	universe.Set(11, 10, true)


	return universe
}
