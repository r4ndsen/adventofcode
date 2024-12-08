package main

import (
	"flag"
	"fmt"
	"github.com/r4ndsen/adventofcode/cast"
	"github.com/r4ndsen/adventofcode/support"
	"log"
	"strings"
)

// var sampleInput = support.InputType("7 6 4 2 1\n1 2 7 8 9\n9 7 6 2 1\n1 3 2 4 5\n8 6 4 4 1\n1 3 6 7 9\n")
var sampleInput = support.InputType("8 12 9 11\n\n")

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
		ans := part2(sampleInput)
		//ans := sample(sampleInput)
		fmt.Println("Output:", ans)
	}
}

func part1(input support.Input) int {
	return sample(input)
}

func part2(input support.Input) int {
	measurements := parseInput(input)

	safe := 0

	for _, inputs := range measurements {
		if isInputSafe(inputs) {
			log.Printf("is safe:\n%v\n", inputs)
			safe++
			continue
		}

		for i := range inputs {
			// input without one index
			test := make([]int, 0)
			for idx := range inputs {
				if i != idx {
					test = append(test, inputs[idx])
				}
			}

			if isInputSafe(test) {
				safe++
				break
			}
		}
	}

	return safe
}

func parseInput(input support.Input) [][]int {
	res := make([][]int, 0)

	for _, l := range input.Lines() {
		if len(strings.TrimSpace(l)) == 0 {
			continue
		}

		res = append(res, cast.LineOfInts(l))
	}

	return res
}

func isInputSafe(inputs []int) bool {
	mode := 0 // 1 == increase -1 == decrease
	for i, m := range inputs {
		if i == 0 {
			continue
		}

		if mode == 0 {
			if m > inputs[i-1] {
				mode = 1
			} else {
				mode = -1
			}
		}

		if m == inputs[i-1] {
			return false
		}

		diff := m - inputs[i-1]

		if diff*mode > 3 {
			return false
		}

		if diff*mode < 0 {
			return false
		}
	}

	return true
}

func sample(input support.Input) int {
	measurements := parseInput(input)

	safe := 0

next:
	for _, inputs := range measurements {
		mode := 0 // 1 == increase -1 == decrease

		for i, m := range inputs {
			if i == 0 {
				continue
			}

			if mode == 0 {
				if m > inputs[i-1] {
					mode = 1
				} else {
					mode = -1
				}
			}

			if m == inputs[i-1] {
				log.Printf("%v unsafe same: %d\n", inputs, m)
				continue next // same as before
			}

			diff := m - inputs[i-1]

			if diff*mode > 3 {
				log.Printf("%v unsafe diff: [%d, %d] %d\n", inputs, inputs[i-1], m, diff*mode)
				continue next
			}

			if diff*mode < 0 {
				log.Printf("%v switching direction\n", inputs)
				continue next
			}
		}
		safe++
	}

	return safe
}
