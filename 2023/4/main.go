package main

import (
	"flag"
	"fmt"
	"github.com/r4ndsen/adventofcode/cast"
	"github.com/r4ndsen/adventofcode/support"
	"regexp"
	"slices"
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
	input := support.InputType(`Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`)

	return part1(input)
}

func part1(input support.Input) int {
	var sum int

	re, _ := regexp.Compile(`\d+`)
	for _, line := range input.Lines() {
		var score int

		split := strings.Split(line, " | ")

		game := re.FindAllString(split[0], -1)
		game = game[1:]

		draw := re.FindAllString(split[1], -1)

		for _, card := range draw {
			if slices.Contains(game, card) {
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}

		sum += score
	}

	return sum
}

func part2(input support.Input) int {
	cards := make(map[int]int)

	re, _ := regexp.Compile(`\d+`)
	for _, line := range input.Lines() {

		split := strings.Split(line, " | ")

		game := re.FindAllString(split[0], -1)

		gameId := cast.ToInt(game[0])
		cards[gameId]++
		game = game[1:]

		draw := re.FindAllString(split[1], -1)

		wins := 0
		for _, card := range draw {
			if slices.Contains(game, card) {
				wins++
			}
		}

		for i := 1; i <= wins; i++ {
			// you have 5 wins, so the next 5 cards gain as many cards as you have tickets for this game
			cards[gameId+i] += cards[gameId]
		}
	}

	var sum int
	for _, amount := range cards {
		sum += amount
	}

	return sum
}
