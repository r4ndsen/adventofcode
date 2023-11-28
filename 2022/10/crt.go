package main

import "strings"

type Crt []*Pixel

func NewCrt() Crt {
	c := make(Crt, 6*40)

	for i := 0; i < len(c); i++ {
		c[i] = NewPixel()
	}

	return c
}

func (c Crt) reset() {
	for _, p := range c {
		p.off()
	}
}

func (c Crt) String() string {
	var s strings.Builder

	for i, p := range c {
		if i > 0 && i%40 == 0 {
			s.WriteByte('\n')
		}
		s.WriteRune(p.value)
	}

	return s.String()
}
