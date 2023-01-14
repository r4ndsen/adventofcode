package main

import (
	"fmt"
	"strings"
)

type Grid struct {
	m                      map[int]map[int]Material
	minX, maxX, minY, maxY int
}

func (g *Grid) addPath(path Path) {
	for i, v := range path {
		if i != 0 {
			previousCoord := path[i-1]
			g.drawPath(previousCoord, v)
		}
	}
}

func (g *Grid) drawPath(from, to Coordinate) {
	for y := from.y; y <= to.y; y++ {
		for x := from.x; x <= to.x; x++ {
			g.addAndAdjustGridDimensions(Coordinate{x: x, y: y}, PathMaterial)
		}
	}

	for y := to.y; y <= from.y; y++ {
		for x := to.x; x <= from.x; x++ {
			g.addAndAdjustGridDimensions(Coordinate{x: x, y: y}, PathMaterial)
		}
	}
}

func (g *Grid) String() string {
	var s strings.Builder

	for y := g.minY; y <= g.maxY; y++ {
		s.WriteString(fmt.Sprintf("%d ", y))
		for x := g.minX; x <= g.maxX; x++ {
			if v, ok := g.m[y][x]; ok {
				s.WriteRune(rune(v))
			} else {
				s.WriteRune(AirMaterial)
			}
		}
		s.WriteRune('\n')
	}

	return s.String()
}

func (g *Grid) free(c Coordinate) bool {
	if c.y > g.maxY {
		return false
	}

	v, ok := g.m[c.y][c.x]
	if !ok || v == AirMaterial {
		return true
	}

	return false
}

func (g *Grid) DropSands(dropPoint Coordinate, amount int) {
	for i := 0; i < amount; i++ {
		g.DropSand(dropPoint)
	}
}

func (g *Grid) DropSand(dropPoint Coordinate) bool {

	down := Coordinate{dropPoint.x, dropPoint.y + 1}
	if g.free(down) {
		return g.DropSand(down)
	}

	downLeft := Coordinate{dropPoint.x - 1, dropPoint.y + 1}
	if g.free(downLeft) {
		return g.DropSand(downLeft)
	}

	downRight := Coordinate{dropPoint.x + 1, dropPoint.y + 1}
	if g.free(downRight) {
		//fmt.Println("right is free", downRight)
		return g.DropSand(downRight)
	}

	if dropPoint.y >= g.maxY {
		return false
	}

	// cannot drop further
	return g.add(dropPoint, SandMaterial)
}

func (g *Grid) add(c Coordinate, value Material) bool {
	if _, ok := g.m[c.y]; !ok {
		g.m[c.y] = make(map[int]Material)
	}

	if v, ok := g.m[c.y][c.x]; ok {
		if v == SandDropoffPointMaterial {
			return false
		}
	}

	g.m[c.y][c.x] = value
	return true
}

func (g *Grid) addAndAdjustGridDimensions(c Coordinate, value Material) {
	if _, ok := g.m[c.y]; !ok {
		g.m[c.y] = make(map[int]Material)
	}

	g.m[c.y][c.x] = value

	if c.x < g.minX {
		g.minX = c.x
	}
	if c.x > g.maxX {
		g.maxX = c.x
	}
	if c.y < g.minY {
		g.minY = c.y
	}
	if c.y > g.maxY {
		g.maxY = c.y
	}
}

func NewGrid() *Grid {
	g := Grid{m: make(map[int]map[int]Material)}

	g.minY = 999999
	g.minX = 999999

	return &g
}
