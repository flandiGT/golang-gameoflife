package util

import (
	"flag"
)

type ArgumentParser struct {
}

func NewArgumentParser() *ArgumentParser {
	return &ArgumentParser{}
}

func (argumentParser *ArgumentParser) Parse() *Arguments {
	const PATTERN_FLAG_NAME = "pattern"
	const DEFAULT_PATTERN = "random"

	const INTERVAL_FLAG_NAME = "interval"
	const DEFAULT_INTERVAL = 500

	const WIDTH_FLAG_NAME = "width"
	const DEFAULT_WIDTH = 160

	const HEIGHT_FLAG_NAME = "height"
	const DEFAULT_HEIGHT = 80

	pattern := flag.String(PATTERN_FLAG_NAME, DEFAULT_PATTERN, "--pattern your_pattern.txt")
	interval := flag.Int(INTERVAL_FLAG_NAME, DEFAULT_INTERVAL, "--interval 500")
	width := flag.Int(WIDTH_FLAG_NAME, DEFAULT_WIDTH, "--width 160")
	height := flag.Int(HEIGHT_FLAG_NAME, DEFAULT_HEIGHT, "--height 80")

	flag.Parse()

	return &Arguments{Pattern:*pattern, Interval:*interval, Width:*width, Height:*height}
}
