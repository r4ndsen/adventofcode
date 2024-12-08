package main

import (
	"flag"
	"fmt"
	"github.com/r4ndsen/adventofcode/support"
)

var sampleInput = support.InputType("....#.....\n.........#\n..........\n..#.......\n.......#..\n..........\n.#..^.....\n........#.\n#.........\n......#...")

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
		ans := sample(sampleInput)
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

func inBounds(x, y int, data [][]rune) bool {
	return x >= 0 && y >= 0 && x < len(data[0]) && y < len(data)
}

func parseInput(input support.Input) int {
	res := 1

	obstacle := '#'
	visited := 'X'
	path := '.'

	m := input.Runes()

	posX, posY, d := getInitialDirection(m)

	m[posY][posX] = visited

	for inBounds(posX, posY, m) {
		futureX := posX + d.X()
		futureY := posY + d.Y()

		// guard left the map
		if !inBounds(futureX, futureY, m) {
			break
		}

		futurField := m[futureY][futureX]

		if futurField == obstacle {
			d = d.TurnRight()
			//fmt.Println("turning right", posY, posX, d)
			continue
		}
		if futurField == path || futurField == visited {
			posX = futureX
			posY = futureY
			m[posY][posX] = visited

			fmt.Println("move to", posY, posX, d)
			if futurField == path {
				res++
			}
			continue
		}
		break
	}

	return res
}

func getInitialDirection(m [][]rune) (posX, posY int, direction support.Direction) {
	for y, row := range m {
		for x, cell := range row {
			posX = x
			posY = y

			if cell == '^' {
				direction = support.North
				return
			} else if cell == 'v' {
				direction = support.South
				return
			} else if cell == '<' {
				direction = support.West
				return
			} else if cell == '>' {
				direction = support.East
				return
			}
		}
	}

	return 0, 0, support.Direction{}
}

func parseInput2(input support.Input) int {
	res := 0

	// used for checking if the current position has been traversed in that direction
	pathDirections := make(map[int]map[int]map[support.Direction]bool)

	obstacle := '#'
	visited := 'X'
	path := '.'

	initMap := input.Runes()
	m := input.Runes()

	posX, posY, d := getInitialDirection(m)

	m[posY][posX] = visited

	for inBounds(posX, posY, m) {
		futureX := posX + d.X()
		futureY := posY + d.Y()

		// guard left the map
		if !inBounds(futureX, futureY, m) {
			break
		}

		futurField := m[futureY][futureX]
		//currentField := m[posY][posX]

		if futurField == obstacle {
			d = d.TurnRight()
			continue
		}

		if futurField == visited {

			if pathDirections[futureX] == nil {
				pathDirections[futureX] = make(map[int]map[support.Direction]bool)
			}

			if pathDirections[futureX][futureY] == nil {
				pathDirections[futureX][futureY] = make(map[support.Direction]bool)
			}

			if pathDirections[futureX][futureY][d] == false {
				pathDirections[futureX][futureY][d] = true
				fmt.Println("placing obstacle at", futureX, futureY)
				res++ // place a new obstacle
			}
		}

		if futurField == path {
			m[posY][posX] = visited
			posX = futureX
			posY = futureY
			continue
		}
		break
	}

	return res
}
