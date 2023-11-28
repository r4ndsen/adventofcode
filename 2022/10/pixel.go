package main

type Pixel struct {
	value rune
}

func (p *Pixel) on() {
	p.value = '#'
}

func (p *Pixel) off() {
	p.value = '.'
}

func NewPixel() *Pixel {
	p := new(Pixel)
	p.off()

	return p
}
