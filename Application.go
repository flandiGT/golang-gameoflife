package main

import (
	"gameoflife/util"
	"gameoflife/renderer"
	"time"
	"gameoflife/factory"
)

func main() {
	arguments := util.NewArgumentParser().Parse()
	universeFactory := factory.SelectUniverseFactory(arguments)

	universe := universeFactory.Build()
	renderer := renderer.NewNcursesRenderer(arguments.Width, arguments.Height)

	universe.SetRenderer(renderer)

	for {
		time.Sleep(time.Duration(arguments.Interval) * time.Millisecond)
		universe.Next()
	}
}
