package main

import (
	"flag"
	"fmt"
	"github.com/r4ndsen/adventofcode/support"
)

var sampleInput = support.InputType("MMMSXXMASM\nMSAMXMSMSA\nAMXSXMAAMM\nMSAMASMSMX\nXMASAMXAMM\nXXAMMXXAMA\nSMSMSASXSS\nSAXAMASAAA\nMAMMMXMMMM\nMXMXAXMASX\n")
var sampleInput2 = support.InputType(".M.S......\n..A..MSMS.\n.M.S.MAA..\n..A.ASMSM.\n.M.S.M....\n..........\nS.S.S.S.S.\n.A.A.A.A..\nM.M.M.M.M.\n..........\n")

var xmas = []byte("XMAS")

func main() {
	var part int
	flag.IntVar(&part, "part", 3, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(support.GetInput())
		fmt.Println("Output:", ans)
	} else if part == 2 {
		ans := part2(support.GetInput())

		fmt.Println("Output:", ans)
	} else {
		ans := sample(sampleInput2)
		fmt.Println("Output:", ans)
	}
}

func part1(input support.Input) int {
	return parseInput(input)
}

func part2(input support.Input) int {
	return parseInput2(input)
}

func sample(input support.Input) int {
	return parseInput2(input)
}

func inBounds(x, y int, data [][]byte) bool {
	return x >= 0 && y >= 0 && x < len(data[0]) && y < len(data)
}

func parseInput(input support.Input) int {
	res := 0

	data := input.Bytes()

	for y, row := range data {
		for x, letter := range row {
			if letter != 'X' {
				continue
			}

		nextdirection:
			for _, d := range support.Directions() {
				currentSearchMultiplyer := 1
				currentSearchIndex := 1
				currentSearchLetter := xmas[currentSearchIndex]

				checkY := y + d.Y()
				checkX := x + d.X()

				for inBounds(checkX, checkY, data) {
					checkLetter := data[checkY][checkX]
					currentSearchMultiplyer++

					if checkLetter != currentSearchLetter {
						continue nextdirection
					}

					currentSearchIndex++

					if currentSearchLetter == 'S' {
						res++
						continue nextdirection
					}

					checkY = y + currentSearchMultiplyer*d.Y()
					checkX = x + currentSearchMultiplyer*d.X()

					currentSearchLetter = xmas[currentSearchIndex]
				}
			}
		}
	}

	return res
}

func parseInput2(input support.Input) int {
	res := 0

	data := input.Bytes()

	checkDiagonalTopLeft := func(x, y int) bool {
		return (data[y-1][x-1] == 'M' && data[y+1][x+1] == 'S') ||
			(data[y-1][x-1] == 'S' && data[y+1][x+1] == 'M')
	}

	checkDiagonalTopRight := func(x, y int) bool {
		return (data[y-1][x+1] == 'M' && data[y+1][x-1] == 'S') ||
			(data[y-1][x+1] == 'S' && data[y+1][x-1] == 'M')
	}

	for y, row := range data {
		if y == 0 || y == len(data)-1 {
			continue
		}

		for x, letter := range row {
			if x == 0 || x == len(row)-1 {
				continue
			}

			if letter != 'A' {
				continue
			}

			if !checkDiagonalTopLeft(x, y) || !checkDiagonalTopRight(x, y) {
				continue
			}

			res++
		}
	}

	return res
}
