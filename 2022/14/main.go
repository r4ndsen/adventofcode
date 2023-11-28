package main

import (
	"fmt"
	"github.com/r4ndsen/adventofcode/support"
	"strings"
)

func main() {
	g := NewGrid()

	sandDropOffPoint := Coordinate{x: 500, y: 0}
	g.addAndAdjustGridDimensions(sandDropOffPoint, SandDropoffPointMaterial)

	for _, line := range support.GetInputFor(14) {
		if len(line) == 0 {
			continue
		}

		var path Path

		for _, pair := range strings.Split(string(line), " -> ") {
			values := strings.Split(pair, ",")
			path = append(path, Coordinate{support.ToInt(values[0]), support.ToInt(values[1])})
		}

		g.addPath(path)
	}

	sandcount := 0

	shallPourMore := true

	for {
		shallPourMore = g.DropSand(sandDropOffPoint)
		if !shallPourMore {
			break
		}
		sandcount++
	}

	fmt.Println(g, "\nsand count with bottomless pit:", sandcount)

	floor := Path{
		Coordinate{x: g.minX - 1000, y: g.maxY + 2},
		Coordinate{x: g.maxX + 1000, y: g.maxY + 2},
	}

	g.addPath(floor)

	for {
		shallPourMore = g.DropSand(sandDropOffPoint)
		if !shallPourMore {
			break
		}
		sandcount++
	}

	fmt.Println("sand count after adding the floor :", sandcount)
}
