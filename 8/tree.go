package main

import (
	"fmt"
)

type Tree struct {
	height      int
	visible     bool
	scenicScore int
}

func (t Tree) String() string {
	return fmt.Sprint(t.height)
}

func (t Tree) Visibility() string {
	if t.visible {
		return "."
	}
	return "X"
}
