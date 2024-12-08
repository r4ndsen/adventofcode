package main

import (
	"flag"
	"fmt"
	"github.com/r4ndsen/adventofcode/support"
)

var sampleInput = support.InputType("")

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
	return parseInput(input)
}

func parseInput(input support.Input) int {
	res := 0

	// todo implement

	return res
}

func parseInput2(input support.Input) int {
	res := 0

	// todo implement

	return res
}
