package main

const (
	Card2 Card = iota
)

type Card rune

func CardFromString(s string) Card {
	return Card(s[0])
}

type Deck struct {
	cards [5]Card
}

func (d *Deck) FiveOfAKind() bool {
	return d.cards[0] == d.cards[1] &&
		d.cards[1] == d.cards[2] &&
		d.cards[2] == d.cards[3] &&
		d.cards[3] == d.cards[4]
}
