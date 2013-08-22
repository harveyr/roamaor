package domain

import (
	"fmt"
	"math/rand"
)

type DiceRoll struct {
	Num int
	Sides int
	Modifier int
}

func (d *DiceRoll) Roll() (roll int) {
	roll = d.Modifier
	for i := 0; i < d.Num; i++ {
		roll += rand.Intn(d.Sides) + 1
	}
	// fmt.Printf("Roll result for %s: %d (random: %d)\n", d, roll, r)
	return
}

func (d *DiceRoll) String() string {
	return fmt.Sprintf("<DiceRoll: %dd%d + %d>", d.Num, d.Sides, d.Modifier)
}

func NewDiceRoll(num int, sides int, mod int) *DiceRoll {
	d := new(DiceRoll)
	d.Num = num
	d.Sides = sides
	d.Modifier = mod
	return d
}
