package main

type Clock struct {
	cycle  int
	X      int
	m      map[int]int
}

func (c *Clock) tick() {
	c.m[c.cycle] = c.X
	c.cycle++ // 1 operation

}

func (c *Clock) noop() {
	c.tick()
}

func (c *Clock) add(x int) {
	c.tick()
	c.X += x
}

func NewClock() *Clock {
	c := Clock{
		X: 1,
		m: make(map[int]int),
	}

	c.tick()

	return &c
}
