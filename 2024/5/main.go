package main

import (
	"flag"
	"fmt"
	"github.com/r4ndsen/adventofcode/cast"
	"github.com/r4ndsen/adventofcode/support"
	"sort"
	"strings"
	"time"
)

var sampleInput = support.InputType("47|53\n97|13\n97|61\n97|47\n75|29\n61|13\n75|53\n29|13\n97|29\n53|29\n61|53\n97|53\n61|29\n47|13\n75|47\n97|75\n47|61\n75|61\n47|29\n75|13\n53|13\n\n75,47,61,53,29\n97,61,53,29,13\n75,29,13\n75,97,47,61,53\n61,13,29\n97,13,75,29,47")

var shouldComeBefore = map[int]map[int]bool{}
var shouldComeAfter = map[int]map[int]bool{}

var updates [][]int

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
	return parseInput2(input)
}

func parseBeforeAndAfter(input support.Input) {

	parsePages := false

	for _, s := range input.Lines() {
		if s == "" {
			parsePages = true
			continue
		}

		if parsePages {
			parts := strings.Split(s, ",")
			ints := make([]int, 0)

			for _, p := range parts {
				ints = append(ints, cast.ToInt(p))
			}
			updates = append(updates, ints)
		} else {
			parts := strings.Split(s, "|")

			left := cast.ToInt(parts[0])
			right := cast.ToInt(parts[1])

			if len(shouldComeBefore[left]) == 0 {
				shouldComeBefore[left] = make(map[int]bool)
			}

			shouldComeBefore[left][right] = true

			if len(shouldComeAfter[right]) == 0 {
				shouldComeAfter[right] = make(map[int]bool)
			}

			shouldComeAfter[right][left] = true
		}
	}
}

func hasConflictBefore(current, other int) bool {
	_, ok := shouldComeAfter[current][other]

	return !ok
}
func hasConflictAfter(current, other int) bool {
	_, ok := shouldComeBefore[current][other]

	return !ok
}

func isInCorrectOrder(update []int) bool {
	for i, v := range update {
		beforeIt := update[:i]
		afterIt := update[i+1:]

		for _, b := range beforeIt {
			if hasConflictBefore(v, b) {
				return false
			}
		}

		for _, a := range afterIt {
			if hasConflictAfter(v, a) {
				return false
			}
		}
	}

	return true
}

func parseInput(input support.Input) int {
	res := 0

	parseBeforeAndAfter(input)

	for _, update := range updates {
		if isInCorrectOrder(update) {
			res += pickMiddleElement(update)
		}
	}

	return res
}

func customSort(arr []int) {
	sort.Slice(arr, func(i, j int) bool {
		a, b := arr[i], arr[j]

		if shouldComeBefore[a] != nil && shouldComeBefore[a][b] {
			return true
		}

		if shouldComeBefore[b] != nil && shouldComeBefore[b][a] {
			return false
		}

		// Default to standard numeric comparison if no precedence rule
		return a < b
	})
}

func pickMiddleElement(arr []int) int {
	middleElement := len(arr) / 2
	return arr[middleElement]
}

func parseInput2(input support.Input) int {
	res := 0

	parseBeforeAndAfter(input)

	start := time.Now()

	for _, update := range updates {
		if isInCorrectOrder(update) {
			continue
		}

		customSort(update)

		res += pickMiddleElement(update)
	}

	fmt.Println("took:", time.Since(start))

	return res
}
