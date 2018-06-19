package factory

import (
	"gameoflife/model"
	"math/rand"
	"time"
)

type RandomUniverseFactory struct {
	universeFactory *UniverseFactory
	width int
	height int
}

func NewRandomUniverseFactory(width int, height int) UniverseFactory {
	return &RandomUniverseFactory{width:width, height:height}
}

func (randomUniverseFactory *RandomUniverseFactory) Build() *model.Universe {
	universe := model.NewUniverse(randomUniverseFactory.width, randomUniverseFactory.height)

	rand.Seed(time.Now().UTC().UnixNano())

	for x:= 0; x < randomUniverseFactory.width; x++ {
		for y := 0; y < randomUniverseFactory.height; y++ {
			universe.Set(x, y, rand.Int() % 3 == 1)
		}
	}

	return universe
}
