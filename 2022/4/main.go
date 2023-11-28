package main

import (
	"fmt"
	"github.com/r4ndsen/adventofcode/support"
	"strconv"
)

type Elf struct {
	minVal, maxVal int
}

func ElfFromRange(r []byte) *Elf {
	for i, v := range r {
		if v == '-' {
			min, _ := strconv.Atoi(string(r[:i]))
			max, _ := strconv.Atoi(string(r[i+1:]))

			return &Elf{min, max}
		}
	}

	return nil
}

func (e *Elf) covers(other *Elf) bool {
	return e.minVal <= other.minVal && e.maxVal >= other.maxVal
}

func (e *Elf) String() string {
	return fmt.Sprintf("%v - %v", e.minVal, e.maxVal)
}

func (e *Elf) overlaps(other *Elf) bool {
	for e1 := e.minVal; e1 <= e.maxVal; e1++ {
		for e2 := other.minVal; e2 <= other.maxVal; e2++ {
			if e1 == e2 {
				return true
			}
		}
	}
	return false
}

func main() {

	sum := 0
	sum2 := 0

	for _, line := range support.GetInputFor(4) {
		for i, v := range line {
			if v != ',' {
				continue
			}

			e1 := ElfFromRange(line[:i])
			e2 := ElfFromRange(line[i+1:])

			if e1.covers(e2) || e2.covers(e1) {
				sum++
			}

			if e1.overlaps(e2) {
				sum2++
			}
			break
		}
	}

	fmt.Println("covers: ", sum, "overlaps:", sum2)
}
