package main

import (
	"flag"
	"fmt"
	"github.com/r4ndsen/adventofcode/cast"
	"github.com/r4ndsen/adventofcode/support"
	"golang.org/x/exp/maps"
	"regexp"
	"strings"
)

var numbers map[string]int
var re *regexp.Regexp

func init() {
	numbers := make(map[string]int)

	for i, v := range []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"} {
		idx := i + 1

		numbers[fmt.Sprintf("%v", idx)] = idx
		numbers[fmt.Sprintf("%v", v)] = idx
	}

	re = regexp.MustCompile(strings.Join(maps.Keys(numbers), "|"))
}

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
		ans := sample(support.InputType("1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet"))
		fmt.Println("Output:", ans)
	}
}

func part1(input support.Input) int {
	ans := 0
	for _, line := range input.Lines() {
		ans += parseLine(line)
	}

	return ans
}

func part2(input support.Input) int {
	var result int
	for _, row := range input.Lines() {
		one, two := firstAndLastDigitWithSpelled(row)
		result += one*10 + two
	}

	return result
}

func firstAndLastDigitWithSpelled(row string) (first int, last int) {
	m := re.FindAllString(row, -1)

	return numbers[m[0]], numbers[m[len(m)-1]]
}

func sample(input support.Input) int {
	ans := 0
	for _, line := range input.Lines() {
		ans += parseLine(line)
	}

	return ans
}

func parseLine(input string) int {
	first, last := firstAndLastDigit(input)

	return first*10 + last
}

func firstAndLastDigit(input string) (first int, last int) {
	re, _ := regexp.Compile("[^1-9]+")
	res := re.ReplaceAllString(input, "")

	return cast.ToInt(res[0]), cast.ToInt(res[len(res)-1])
}
