package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

var priorityByChar = map[byte]int{}

func main() {
	i := 1
	for r := 'a'; r <= 'z'; r++ {
		priorityByChar[byte(r)] = i
		i++
	}
	for r := 'A'; r <= 'Z'; r++ {
		priorityByChar[byte(r)] = i
		i++
	}

	f, err := os.Open("moves.txt")
	if err != nil {
		log.Fatal(err)
	}

	r := bufio.NewReader(f)

	sum := 0

	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}

		sum += GetPriorityFromInput(line)
	}

	fmt.Println(sum)
}

func GetPriorityFromInput(input []byte) int {
	m1 := make(map[byte]int)
	m2 := make(map[byte]int)

	for _, v := range input[:len(input)/2] {
		m1[v]++
	}

	for _, v := range input[len(input)/2:] {
		m2[v]++
	}

	for char := range m1 {
		if _, ok := m2[char]; ok {
			return priorityByChar[char]
		}
	}

	log.Fatal("not found" + string(input))
	return 0
}
