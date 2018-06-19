package factory

import (
	"gameoflife/model"
	"io/ioutil"
	"strings"
	"strconv"
	"fmt"
)

type TextFileUniverseFactory struct {
	UniverseFactory
	fileName string
}

func NewTextFileUniverseFactory(fileName string) UniverseFactory {
	return TextFileUniverseFactory{fileName:fileName}
}

func (reader TextFileUniverseFactory) Build() *model.Universe {
	b, err := ioutil.ReadFile(reader.fileName)
	if err != nil {
		fmt.Print(err)
	}

	content := string(b)
	lines := strings.Split(content, "\n")

	width, err := strconv.Atoi(lines[0])
	if err != nil {
		fmt.Print(err)
	}

	height, err := strconv.Atoi(lines[1])
	if err != nil {
		fmt.Print(err)
	}

	universe := model.NewUniverse(width, height)
	for lineIndex := 2; lineIndex < len(lines); lineIndex++ {
		line := lines[lineIndex]

		for charIndex := 0; charIndex < len(line); charIndex++ {
			if line[charIndex] == '0' {
				universe.Set(charIndex, lineIndex - 2, true)
			}
		}
	}

	return universe
}
