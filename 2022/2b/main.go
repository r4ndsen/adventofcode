package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type Shape int

func (s Shape) win() int {
	return 6
}

func (s Shape) draw() int {
	return 3
}

func (s Shape) scoreByPick() int {
	scoreByPick, _ := scoreByPick[s]
	return scoreByPick
}

func (s Shape) scoreAgainst(opponent Shape) int {
	// draw
	if s == opponent {
		return s.draw() + s.scoreByPick()
	}

	loser, _ := winsAgainst[s]

	// win
	if loser == opponent {
		return s.win() + s.scoreByPick()
	}

	// loss
	return s.scoreByPick()
}

func (s Shape) nemesis() Shape {
	for nemesis, me := range winsAgainst {
		if s == me {
			return nemesis
		}
	}

	log.Fatal("invalid shape: has no nemesis")
	return Shape(0)
}

func (s Shape) favoriteEnemy() Shape {
	ememy, _ := winsAgainst[s]
	return ememy
}

var (
	ROCK     = Shape(1)
	PAPER    = Shape(2)
	SCISSORS = Shape(3)
)

func ShapeFromRune(input rune) Shape {
	switch input {
	case 'A', 'X':
		return ROCK
	case 'B', 'Y':
		return PAPER
	case 'C', 'Z':
		return SCISSORS
	default:
		log.Fatalf("invalid shape: %v", input)
		return Shape(0)
	}
}

var winsAgainst = map[Shape]Shape{
	ROCK:     SCISSORS,
	PAPER:    ROCK,
	SCISSORS: PAPER,
}

var scoreByPick = map[Shape]int{
	ROCK:     1,
	PAPER:    2,
	SCISSORS: 3,
}

func main() {
	f, err := os.Open("moves.txt")
	if err != nil {
		log.Fatal(err)
	}

	r := bufio.NewReader(f)

	sumScore := 0

	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}

		if len(line) == 0 {
			continue
		}

		opponentPick := ShapeFromRune(rune(line[:1][0]))
		action := line[2:3][0]

		yourPick := opponentPick

		switch action {
		case 'X':
			yourPick = opponentPick.favoriteEnemy()
		case 'Z':
			yourPick = opponentPick.nemesis()
		}

		sumScore += yourPick.scoreAgainst(opponentPick)
	}
	fmt.Println(sumScore)
}
