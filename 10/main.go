package main

import (
	"fmt"
	"github.com/r4ndsen/adventofcode/support"
	"strconv"
)

func main() {

	clock := NewClock()

	for _, v := range support.GetInputFor(10) {
		if string(v[:4]) == "addx" {
			val, err := strconv.Atoi(string(v[5:]))
			support.Check(err)

			clock.add(val)
		}

		clock.noop()
	}

	checkIndexes := []int{20, 60, 100, 140, 180, 220}
	sum := 0
	for _, i := range checkIndexes {
		signalStrength := i * clock.m[i-1]
		fmt.Printf("signal strength %v at cycle %d: %d\n", i, signalStrength, clock.m[i-1])
		sum += signalStrength
	}
	fmt.Printf("Signal Strength: %v\n", sum)

	crt := NewCrt()

	offset := -40

	for i := 0; i < len(clock.m)-1; i++ {
		if i%40 == 0 {
			offset += 40
		}
		crt[i].off()

		if clock.m[i]+offset >= i-1 && clock.m[i]+offset <= i+1 {
			crt[i].on()
		}
	}

	fmt.Println(crt)
}
