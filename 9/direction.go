package main

import "fmt"

type Vector struct {
	x, y int
}

func (v Vector) String() string {
	return fmt.Sprintf("[ %v | %v ]", v.x, v.y)
}

func (v Vector) direction() Vector {
	var newVector Vector

	if v.x > 0 {
		newVector.x = 1
	} else if v.x < 0 {
		newVector.x = -1
	}

	if v.y > 0 {
		newVector.y = 1
	} else if v.y < 0 {
		newVector.y = -1
	}

	return newVector
}

func NewVector(v1, v2 *Vector) Vector {
	return Vector{
		x: v2.x - v1.x,
		y: v2.y - v1.y,
	}
}
