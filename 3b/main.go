package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

var priorityByChar = map[byte]int{}

type ElfGroup struct {
	elves            []*Elf
	groupIdentifyers []byte
	solved           bool
}

type Elf []byte

func (elf *Elf) uniqueItems() []byte {
	result := make([]byte, 0)
	cache := make(map[byte]bool)

	for _, item := range *elf {
		if cache[item] {
			continue
		}

		result = append(result, item)
		cache[item] = true
	}

	return result
}

func (elf *Elf) hasItem(item byte) bool {
	for _, i := range *elf {
		if i == item {
			return true
		}
	}

	return false
}

func (elf *Elf) String() string {
	return string(*elf)
}

func NewElf(data []byte) *Elf {
	e := Elf(data)

	return &e
}

func (elfGroup *ElfGroup) GroupName() string {
	return string(elfGroup.groupIdentifyers)
}

func (elfGroup *ElfGroup) resolveCommonItems() {
	for _, item := range elfGroup.elves[0].uniqueItems() {
		if elfGroup.elves[1].hasItem(item) && elfGroup.elves[2].hasItem(item) {
			elfGroup.groupIdentifyers = append(elfGroup.groupIdentifyers, item)
		}
	}

	elfGroup.solved = len(elfGroup.groupIdentifyers) == 1
}

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

	lineIndex := 0
	sum := 0
	group := new(ElfGroup)
	for {
		line, err := r.ReadBytes(byte('\n'))

		if err == io.EOF {
			break
		}

		if len(line) == 0 {
			break
		}

		fmt.Println(string(line[:len(line)-1]))

		group.elves = append(group.elves, NewElf(line[:len(line)-1]))

		lineIndex++

		if lineIndex%3 == 0 {
			group.resolveCommonItems()
			sum += group.Score()

			group = new(ElfGroup)
		}
	}

	fmt.Println(sum)
}

func (elfGroup *ElfGroup) Score() int {
	return priorityByChar[elfGroup.groupIdentifyers[0]]
}
