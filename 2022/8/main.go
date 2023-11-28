package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {

	f, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	r := bufio.NewReader(f)

	grid := make(Grid, 0)

	lineNr := 0

	for {
		line, err := r.ReadBytes(byte('\n'))

		if err == io.EOF {
			break
		}

		if len(line) == 0 {
			break
		}

		grid = append(grid, make([]*Tree, len(line)-1))

		for idx, v := range line[:len(line)-1] {
			height, _ := strconv.Atoi(string([]byte{v}))
			grid[lineNr][idx] = &Tree{height: height}
		}

		lineNr++
	}

	//grid.determineVisibilities()
	grid.determineScenicScores()

	//grid.determineScenicScoreForTreeAt(2, 3)

	//fmt.Println(grid.MapScenicScores())

	//fmt.Println(grid.String())
	//fmt.Println(grid.VisibleTrees())
	//fmt.Println(grid.CountVisible())
	fmt.Println(grid.BestTreeForBuildingTheTreeHouse().scenicScore)
}
