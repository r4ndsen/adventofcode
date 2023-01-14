package main

import (
	"bufio"
	"fmt"
	"github.com/r4ndsen/adventofcode/support"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Stack struct {
	items []*Crate
}

func (s *Stack) addCrate(c *Crate) {
	s.addToBottom(c)
}

func (s *Stack) addToBottom(c *Crate) {
	s.items = append([]*Crate{c}, s.items...)
}

func (s *Stack) popItems(amount int) []*Crate {
	itemCount := len(s.items)

	bottom := s.items[:itemCount-amount]
	top := s.items[itemCount-amount:]

	s.items = bottom

	return top
}

type Crate byte

func (c *Stack) String() string {
	var s strings.Builder

	for i := 0; i < len(c.items); i++ {
		s.WriteString(c.items[i].String())
	}

	return s.String()
}

func (c *Crate) String() string {
	return string(*c)
}

func (s *Stack) push(c ...*Crate) {
	s.items = append(s.items, c...)
}

func (s *Stack) pop() *Crate {
	c := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return c
}

func (s *Stack) moveItemsToStack(amount int, other *Stack) {
	crates := s.popItems(amount)
	other.push(crates...)
}

func main() {

	f, err := os.Open("start.txt")
	if err != nil {
		log.Fatal(err)
	}

	r := bufio.NewReader(f)
	stacks := MakeStacks(9)

	for {
		line, err := r.ReadBytes(byte('\n'))

		if err == io.EOF {
			break
		}

		if len(line) == 0 {
			break
		}

		line = support.Trim('\n', line)

		for i, v := range line {
			if ((i + 3) % 4) == 0 {
				if v == ' ' {
					continue
				}
				c := Crate(v)
				stacks[StackFromIndex(i)].addCrate(&c)
			}
		}
	}

	f, err = os.Open("moves.txt")
	if err != nil {
		log.Fatal(err)
	}

	r = bufio.NewReader(f)
	re := regexp.MustCompile(`move (\d+) from (\d) to (\d)`)

	for {
		line, err := r.ReadBytes(byte('\n'))

		if err == io.EOF {
			break
		}

		if len(line) == 0 {
			break
		}

		res := re.FindAllStringSubmatch(string(line), -1)

		amount, _ := strconv.Atoi(res[0][1])
		from, _ := strconv.Atoi(res[0][2])
		to, _ := strconv.Atoi(res[0][3])

		fromStack := stacks[from]
		toStack := stacks[to]

		fmt.Printf("move %v from %v to %v\n", amount, from, to)

		fromStack.moveItemsToStack(amount, toStack)
	}

	for i := 1; i <= len(stacks); i++ {
		fmt.Print(stacks[i].pop())
	}
}

func MakeStacks(amount int) map[int]*Stack {
	stacks := make(map[int]*Stack)

	for i := 1; i <= amount; i++ {
		stacks[i] = new(Stack)
	}

	return stacks
}

func StackFromIndex(i int) int {
	return 1 + i/4
}
