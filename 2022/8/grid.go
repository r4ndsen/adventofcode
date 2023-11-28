package main

import (
	"fmt"
	"sort"
	"strings"
)

type Grid [][]*Tree

func (g Grid) CountVisible() int {
	res := 0

	for x, line := range g {
		for y := range line {
			if g[x][y].visible {
				res++
			}
		}
	}

	return res
}

func (g Grid) String() string {

	var s strings.Builder

	for _, line := range g {
		for _, v := range line {
			s.WriteString(v.String())
		}
		s.WriteString("\n")
	}

	return s.String()
}

func (g Grid) VisibleTrees() string {

	var s strings.Builder

	for _, line := range g {
		for _, v := range line {
			s.WriteString(v.Visibility())
		}
		s.WriteString("\n")
	}

	return s.String()
}

func (g Grid) MapScenicScores() string {

	var s strings.Builder

	for _, line := range g {
		for _, t := range line {
			s.WriteString(fmt.Sprintf("%v", t.scenicScore))
		}
		s.WriteString("\n")
	}

	return s.String()
}

func (g Grid) determineVisibilities() {
	for rowIdx, line := range g {
		for colIdx, tree := range line {
			tree.visible = g.determineVisibilityForTreeAt(rowIdx, colIdx)
		}
	}
}

func (g Grid) determineScenicScores() {
	for rowIdx, line := range g {
		for colIdx, tree := range line {
			tree.scenicScore = g.determineScenicScoreForTreeAt(rowIdx, colIdx)
		}
	}
}

func (g Grid) determineVisibilityForTreeAt(x, y int) bool {
	if x == 0 || y == 0 || x == len(g)-1 || y == len(g)-1 {
		return true
	}

	t := g[y][x]

	visibleFromLeft := func(x, y int) bool {
		for i := x - 1; i >= 0; i-- {
			if g[y][i].height >= t.height {
				return false
			}
		}
		return true
	}

	visibleFromRight := func(x, y int) bool {
		for i := x + 1; i < len(g); i++ {
			if g[y][i].height >= t.height {
				return false
			}
		}
		return true
	}

	visibleFromTop := func(x, y int) bool {
		for i := y - 1; i >= 0; i-- {
			if g[i][x].height >= t.height {
				return false
			}
		}
		return true
	}

	visibleFromBottom := func(x, y int) bool {
		for i := y + 1; i < len(g); i++ {
			if g[i][x].height >= t.height {
				return false
			}
		}
		return true
	}

	return visibleFromTop(x, y) || visibleFromRight(x, y) || visibleFromBottom(x, y) || visibleFromLeft(x, y)
}

func (g Grid) determineScenicScoreForTreeAt(x, y int) int {
	t := g[y][x]

	//fmt.Println("tree at ", x, y, t.height)

	treesTop := func(x, y int) int {
		distance := 0
		for i := x - 1; i >= 0; i-- {
			distance++
			if g[y][i].height >= t.height {
				break
			}
		}
		return distance
	}

	treesRight := func(x, y int) int {
		distance := 0
		for i := x + 1; i < len(g); i++ {
			distance++
			if g[y][i].height >= t.height {
				break
			}
		}
		return distance
	}

	treesBottom := func(x, y int) int {
		distance := 0
		for i := y - 1; i >= 0; i-- {
			distance++
			if g[i][x].height >= t.height {
				break
			}
		}
		return distance
	}

	treesLeft := func(x, y int) int {
		distance := 0
		for i := y + 1; i < len(g); i++ {
			distance++
			if g[i][x].height >= t.height {
				break
			}
		}
		return distance
	}

	//fmt.Println(treesBottom(x, y), treesRight(x, y), treesLeft(x, y), treesTop(x, y))
	return treesBottom(x, y) * treesRight(x, y) * treesLeft(x, y) * treesTop(x, y)
}

func (g Grid) BestTreeForBuildingTheTreeHouse() *Tree {
	trees := g.Trees()

	sort.SliceStable(trees, func(i, j int) bool {
		return trees[i].scenicScore > trees[j].scenicScore
	})

	return trees[0]
}

func (g Grid) Trees() []*Tree {
	trees := make([]*Tree, 0)

	for _, row := range g {
		for _, t := range row {
			trees = append(trees, t)
		}
	}

	return trees
}
