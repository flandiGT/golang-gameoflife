package main

import (
	"time"
	"gameoflife/renderer"
	"gameoflife/pattern"
)

func main() {
	const width = 230
	const height = 60

	universe := pattern.CreateRandomUniverse(width, height)
	renderer := renderer.NewNcursesRenderer(width, height)

	universe.SetRenderer(renderer)

	for {
		time.Sleep(200 * time.Millisecond)
		universe.Next()
	}
}
