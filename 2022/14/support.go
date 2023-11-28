package main

type Coordinate struct {
	x, y int
}

type Material rune

type Path []Coordinate

const (
	SandMaterial             = 'o'
	PathMaterial             = '#'
	AirMaterial              = ' '
	SandDropoffPointMaterial = 'x'
)
