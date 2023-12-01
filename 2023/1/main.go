package main

import (
	"flag"
	"fmt"
	"github.com/r4ndsen/adventofcode/cast"
	"github.com/r4ndsen/adventofcode/support"
	"golang.org/x/exp/maps"
	"math"
	"regexp"
	"strings"
)

var numbers map[string]int

func init() {
	numbers := make(map[string]int)

	for i, v := range []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"} {
		idx := i + 1

		numbers[fmt.Sprintf("%v", idx)] = idx
		numbers[fmt.Sprintf("%v", v)] = idx
	}
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
		ans := sample([][]byte{
			[]byte("1abc2"),
			[]byte("pqr3stu8vwx"),
			[]byte("a1b2c3d4e5f"),
			[]byte("treb7uchet"),
		})
		fmt.Println("Output:", ans)
	}
}

func part1(input [][]byte) int {
	ans := 0
	for _, line := range input {
		ans += parseLine(line)
	}

	return ans
}

func part2(input [][]byte) int {
	reString := strings.Join(maps.Keys(numbers), "|")
	re := regexp.MustCompile(reString)

	var result int
	for _, row := range input {
		one, two := firstAndLastDigitWithSpelled(string(row), re)
		result += one*10 + two
	}

	return result
}

func firstAndLastDigitWithSpelled(row string, re *regexp.Regexp) (first int, last int) {
	lastIdx := math.MinInt

	first = numbers[re.FindString(row)]

	for k, v := range numbers {
		if found := strings.LastIndex(row, k); found != -1 && found > lastIdx {
			lastIdx = found
			last = v
		}
	}

	return
}

func sample(input [][]byte) int {
	ans := 0
	for _, line := range input {
		ans += parseLine(line)
	}

	return ans
}

func parseLine(input []byte) int {
	first, last := firstAndLastDigit(input)

	return first*10 + last
}

func firstAndLastDigit(input []byte) (first int, last int) {
	re, _ := regexp.Compile("[^1-9]+")
	res := re.ReplaceAll(input, []byte(""))

	return cast.ToInt(res[0]), cast.ToInt(res[len(res)-1])
}
