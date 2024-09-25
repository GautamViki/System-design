package dice

import "math/rand"

type Dice struct {
	Sides int
}

func NewDice(sides int) *Dice {
	return &Dice{
		Sides: sides,
	}
}

func (d *Dice) Roll() int {
	return rand.Intn(6-1) + 1
}
