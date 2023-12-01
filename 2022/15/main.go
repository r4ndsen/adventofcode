package main

import (
	"fmt"
	"github.com/r4ndsen/adventofcode/cast"
	"github.com/r4ndsen/adventofcode/support"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)

	var sensors []Sensor

	for _, line := range support.GetInputFor(15) {
		if len(line) == 0 {
			continue
		}

		res := re.FindAllStringSubmatch(string(line), -1)

		sensor := Sensor{
			position: Coordinate{
				cast.ToInt(res[0][1]),
				cast.ToInt(res[0][2]),
			},
			beacon: Coordinate{
				cast.ToInt(res[0][3]),
				cast.ToInt(res[0][4]),
			},
		}

		sensors = append(sensors, sensor)
	}

	coveredPositions := make(map[int]bool, 0)

	var left, right int

	for _, s := range sensors {
		coveredRadius := s.ManhattenDistance(s.beacon)

		if s.position.x-coveredRadius < left {
			left = s.position.x - coveredRadius
		}
		if s.position.x+coveredRadius > right {
			right = s.position.x + coveredRadius
		}
	}

positionLoop:
	for i := left; i < right; i++ {
		testPosition := Coordinate{i, 2000000}

		for _, s := range sensors {
			if s.Covers(testPosition) {
				coveredPositions[i] = true
				continue positionLoop
			}
		}
	}

	for _, s := range sensors {
		if s.beacon.y == 2000000 {
			delete(coveredPositions, s.beacon.x)
		}
	}

	fmt.Println("we have", len(coveredPositions), "covered spots")

}
