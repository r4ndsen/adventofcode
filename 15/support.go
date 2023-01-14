package main

import (
	"fmt"
)

type Coordinate struct {
	x, y int
}

func (c Coordinate) TuningFrequency() int {
	return 4000000*c.x + c.y
}

func (c Coordinate) String() string {
	return fmt.Sprintf("%v,%v", c.x, c.y)
}

type Sensor struct {
	position Coordinate
	beacon   Coordinate
}

func (s Sensor) Borders() []Coordinate {
	distance := s.ManhattenDistance(s.beacon)

	res := make([]Coordinate, 0)

	// top right edge
	x, y := s.position.x, s.position.y+distance+1
	for y > s.position.y {
		res = append(res, Coordinate{x, y})
		x++
		y--
	}

	// bottom right edge
	x, y = s.position.x+distance+1, s.position.y
	for x > s.position.x {
		// top right edge
		res = append(res, Coordinate{x, y})
		x--
		y--
	}

	// bottom left edge
	x, y = s.position.x, s.position.y-distance-1
	for y < s.position.y {
		// top right edge
		res = append(res, Coordinate{x, y})
		x--
		y++
	}

	// top left edge
	x, y = s.position.x-distance-1, s.position.y
	for x < s.position.x {
		// top right edge
		res = append(res, Coordinate{x, y})
		x++
		y++
	}

	return res
}

func (s Sensor) ManhattenDistance(c Coordinate) int {
	x := c.x - s.position.x
	if x < 0 {
		x *= -1
	}

	y := c.y - s.position.y
	if y < 0 {
		y *= -1
	}

	return x + y
}

func (s Sensor) Covers(c Coordinate) bool {
	return s.ManhattenDistance(c) <= s.ManhattenDistance(s.beacon)
}
