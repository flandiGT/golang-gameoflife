package main

import (
	"gameoflife/core"
	"time"
)

func main() {
	universe := core.CreateRandomUniverse(100, 100)
	renderer := core.NewNcursesRenderer()

	for {
		renderer.Render(universe)
		time.Sleep(1 * time.Millisecond)
		universe.Next()
	}
}
