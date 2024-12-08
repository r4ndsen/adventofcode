package main

import (
	"flag"
	"fmt"
	"github.com/r4ndsen/adventofcode/cast"
	"github.com/r4ndsen/adventofcode/support"
	"regexp"
)

var sampleInput = support.InputType("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))")
var sampleInput2 = support.InputType("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))")

var match = regexp.MustCompile(`mul\((\d+),(\d+)\)`).FindAllStringSubmatch
var match2 = regexp.MustCompile(`don't\(\)|do\(\)|mul\((\d+),(\d+)\)`).FindAllStringSubmatch

const (
	dont = `don't()`
	do   = `do()`
)

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
	calculations := match(input.String(), -1)

	res := 0
	for _, c := range calculations {
		res += cast.ToInt(c[1]) * cast.ToInt(c[2])
	}

	return res
}

func parseInput2(input support.Input) int {
	instructions := match2(input.String(), -1)

	var pick = true

	isPickInstruction := func(s string) bool {
		return s == dont || s == do
	}

	res := 0
	for _, c := range instructions {
		if isPickInstruction(c[0]) {
			pick = c[0] == do
			continue
		}

		if pick {
			res += cast.ToInt(c[1]) * cast.ToInt(c[2])
		}
	}

	return res
}
