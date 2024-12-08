package main

import (
	"flag"
	"fmt"
	"github.com/r4ndsen/adventofcode/cast"
	"github.com/r4ndsen/adventofcode/support"
	"math"
	"regexp"
	"slices"
	"strings"
)

var match = regexp.MustCompile(`^(\d+)\s+(\d+)$`).FindAllStringSubmatch

func main() {

	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(support.GetInput())
		_ = support.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else if part == 2 {
		ans := part2(support.GetInput())
		_ = support.CopyToClipboard(fmt.Sprintf("%v", ans))

		fmt.Println("Output:", ans)
	} else {
		ans := sample(support.InputType("3   4\n4   3\n2   5\n1   3\n3   9\n3   3\n"))
		fmt.Println("Output:", ans)
	}
}

func part1(input support.Input) int {
	locations := parseInput(input)

	dis := 0
	for i, left := range locations[0] {
		// cast to int
		dis += cast.ToInt(math.Abs(float64(locations[1][i] - left)))
	}

	return dis
}

func part2(input support.Input) int {
	locations := parseInput(input)

	res := 0
	for _, left := range locations[0] {
		cnt := 0
		for _, right := range locations[1] {
			if left == right {
				cnt++
			}
		}

		res += left * cnt
	}

	return res
}

func parseInput(input support.Input) [][]int {
	locations := make([][]int, 2)

	for _, l := range input.Lines() {
		if len(strings.TrimSpace(l)) == 0 {
			continue
		}

		matches := match(l, -1)[0][1:]

		for i, s := range matches {
			locations[i] = append(locations[i], cast.ToInt(s))
		}
	}

	slices.Sort(locations[0])
	slices.Sort(locations[1])

	return locations
}

func sample(input support.Input) int {
	locations := parseInput(input)

	dis := 0
	for i, left := range locations[0] {
		dis += cast.ToInt(math.Abs(float64(locations[1][i] - left)))
	}

	return dis
}
