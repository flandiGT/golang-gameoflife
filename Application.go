package main

import (
	"time"
	"gameoflife/pattern"
	"gameoflife/renderer"
)

func main() {
	universe := pattern.CreateRandomUniverse(230, 60)
	renderer := renderer.NewNcursesRenderer()

	for {
		renderer.Render(universe)
		time.Sleep(1 * time.Millisecond)
		universe.Next()
	}
}
