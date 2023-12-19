package main

import (
	"flag"
	"fmt"
	"github.com/r4ndsen/adventofcode/support"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(support.GetInput())
		support.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else if part == 2 {
		ans := part2(support.GetInput())
		support.CopyToClipboard(fmt.Sprintf("%v", ans))

		fmt.Println("Output:", ans)
	} else {
		ans := sample()
		fmt.Println("Output:", ans)
	}
}

func sample() int {
	input := support.InputType(`32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`)

	return part1(input)
}

func part1(input support.Input) int {

	return 0
}

func part2(input support.Input) int {

	return 1
}
