package factory

import (
	"gameoflife/model"
	"gameoflife/util"
)

type UniverseFactory interface {
	Build() *model.Universe
}

func SelectUniverseFactory(arguments *util.Arguments) UniverseFactory {
	if arguments.Pattern == "random" {
		return NewRandomUniverseFactory(arguments.Width, arguments.Height)
	}

	if &arguments.Pattern == nil || len(arguments.Pattern) == 0 {
		return NewRandomUniverseFactory(arguments.Width, arguments.Height)
	}

	return NewTextFileUniverseFactory(arguments.Pattern)
}
