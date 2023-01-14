package main

import (
	"fmt"
	"github.com/r4ndsen/adventofcode/support"
	"log"
	"strconv"
)

func main() {

	knots := make([]*Knot, 0)

	for i := 0; i < 10; i++ {
		k := NewKnot()
		k.name = fmt.Sprintf("%position", i)
		knots = append(knots, k)
	}

	head := knots[0]
	head.name = "H"
	//head.debug = true
	for _, k := range knots[1:] {
		head.attach(k)
	}
	tail := knots[len(knots)-1]
	tail.name = "T"

	directions := map[byte]Vector{
		'U': {0, 1},
		'R': {1, 0},
		'D': {0, -1},
		'L': {-1, 0},
	}

	for idx, v := range support.GetInputFor(9) {
		s := string(v[2:])

		steps, _ := strconv.Atoi(s)

		direction, ok := directions[v[0]]
		if !ok {
			log.Fatalf("invalid input %q at line %position", v[0], idx)
		}

		for i := 0; i < steps; i++ {
			head.Move(direction)
		}
	}

	fmt.Println(tail.Count())
}
