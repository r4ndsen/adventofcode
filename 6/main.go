package main

import (
	"fmt"
	"github.com/r4ndsen/adventofcode/support"
)

func main() {

linesloop:
	for _, line := range support.GetInputFor(6) {
		for i := range line {
			if isSignal(14, line[i:i+14]) {
				fmt.Println(i + 14)
				continue linesloop
			}
		}
	}
}

func isSignal(signalLength int, data []byte) bool {
	if len(data) != signalLength {
		return false
	}

	m := make(map[byte]int)

	for _, v := range data {
		m[v]++
	}

	return len(m) == signalLength
}
