package main

import (
	"fmt"
)

type Boundary struct {
	Dest, Source, Range int
	conversionOffset    int
	maxSource           int
	Name                string
}

func (b Boundary) Hit(i int) bool {
	return i >= b.Source && i <= b.maxSource
}

func (b Boundary) Convert(i int) int {
	if !b.Hit(i) {
		//fmt.Println("not hit", i, "with", b.String(), "to", i)
		return i
	}

	//fmt.Println("converting", i, "with", b.String(), "to", i+b.conversionOffset)

	return i + b.conversionOffset
}

func (b Boundary) String() string {
	return fmt.Sprintf("%s: [%d-%d] => [%d-%d]", b.Name, b.Source, b.Source+b.Range, b.Dest, b.Dest+b.Range)
}

func NewBoundary(name string, d, s, r int) *Boundary {
	return &Boundary{
		Name:             name,
		Dest:             d,
		Source:           s,
		Range:            r,
		conversionOffset: d - s,
		maxSource:        s + r,
	}
}
