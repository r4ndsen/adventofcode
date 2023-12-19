package main

import (
	"flag"
	"fmt"
	"github.com/r4ndsen/adventofcode/cast"
	"github.com/r4ndsen/adventofcode/support"
	"regexp"
	"strings"
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
	input := support.InputType(`Time:      7  15   30
Distance:  9  40  200`)

	return part2(input)
}

func part1(input support.Input) int {
	re := regexp.MustCompile(`\d+`)
	lines := input.Lines()

	times := re.FindAllString(lines[0], -1)
	distances := re.FindAllString(lines[1], -1)

	result := 0

	for i, time := range times {
		r := Race{Time: cast.ToInt(time), Distance: cast.ToInt(distances[i])}

		winningConfigurations := r.WinningConfigurations()

		if winningConfigurations > 1 {
			if result == 0 {
				result = winningConfigurations
			} else {
				result *= winningConfigurations
			}
		}
	}

	return result
}

func part2(input support.Input) int {

	re := regexp.MustCompile(`\d+`)
	lines := input.Lines()

	times := re.FindAllString(lines[0], -1)
	distances := re.FindAllString(lines[1], -1)

	time := strings.Join(times, "")
	distance := strings.Join(distances, "")

	r := Race{Time: cast.ToInt(time), Distance: cast.ToInt(distance)}

	return r.WinningConfigurations()
}
