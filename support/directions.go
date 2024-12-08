package support

import "log"

type Direction struct {
	icon string
	x    int
	y    int
}

var (
	East      = Direction{"➡️", 1, 0}
	SouthEast = Direction{"↘️", 1, 1}
	South     = Direction{"⬇️", 0, 1}
	SouthWest = Direction{"↙️", -1, 1}
	West      = Direction{"⬅️", -1, 0}
	NorthWest = Direction{"↖️", -1, -1}
	North     = Direction{"⬆️", 0, -1}
	NorthEast = Direction{"↗️", 1, -1}
)

var directions = [8]Direction{
	East,
	SouthEast,
	South,
	SouthWest,
	West,
	NorthWest,
	North,
	NorthEast,
}

func (d Direction) TurnRight() Direction {
	for i, direction := range directions {
		if direction == d {
			return directions[(i+2)%len(directions)]
		}
	}
	log.Fatal("something went terribly wrong")
	return Direction{}
}

func Directions() [8]Direction {
	return directions
}

func (d Direction) X() int {
	return d.x
}

func (d Direction) Y() int {
	return d.y
}

func (d Direction) String() string {
	return d.icon
}
