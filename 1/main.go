package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {

	f, err := os.Open("moves.txt")
	if err != nil {
		log.Fatal(err)
	}

	r := bufio.NewReader(f)

	elves := make(map[int]int)

	elfIndex := 0
	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}

		if len(line) == 0 {
			elfIndex++
			continue
		}

		if calories, err := strconv.Atoi(string(line)); err == nil {
			elves[elfIndex] += calories
		}
	}

	var values []int
	for _, val := range elves {
		values = append(values, val)
	}

	sort.Ints(values)

	sum := 0
	for _, val := range values[len(values)-3:] {
		sum += val
	}

	fmt.Println(sum)
}
