package main

import "fmt"

type Knot struct {
	name     string
	position *Vector
	tail     *Knot
	visited  map[int]map[int]bool
	debug    bool
}

func (k *Knot) registerField() {
	_, ok := k.visited[k.position.x]
	if !ok {
		k.visited[k.position.x] = make(map[int]bool)
	}

	k.visited[k.position.x][k.position.y] = true
}

func (k *Knot) Count() int {
	res := 0
	for i := range k.visited {
		res += len(k.visited[i])
	}

	return res
}

func (k *Knot) follow(other *Knot) {
	if k.hasContact(other) {
		return
	}

	distance := NewVector(k.position, other.position)
	k.Move(distance.direction())
}

func (k *Knot) Move(direction Vector) {

	if k.debug {
		fmt.Printf("%s move from %s to ", k.name, k.position)
	}

	k.position.x += direction.x
	k.position.y += direction.y

	if k.debug {
		fmt.Println(k.position)
	}
	k.registerField()
	if k.tail != nil {
		k.tail.follow(k)
	}
}

func (k *Knot) attach(other *Knot) {
	if k.tail == nil {
		k.tail = other
		return
	}

	k.tail.attach(other)
}

func (k *Knot) hasContact(other *Knot) bool {
	if other.position.x-k.position.x > 1 || other.position.x-k.position.x < -1 {
		return false
	}

	if other.position.y-k.position.y > 1 || other.position.y-k.position.y < -1 {
		return false
	}

	return true
}

func NewKnot() *Knot {
	k := Knot{
		visited:  make(map[int]map[int]bool),
		position: &Vector{},
	}
	k.registerField()

	return &k
}
