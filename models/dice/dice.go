package dice

import (
	"fmt"
	"math/rand"
)

type DiceRoll struct {
	Num uint8
	Sides uint8
	Modifier uint8
}

func (d *DiceRoll) Roll() (roll uint16) {
	roll = uint16(d.Modifier)
	for i := 0; i < int(d.Num); i++ {
		roll += uint16(rand.Intn(int(d.Sides)) + 1)
	}
	r := rand.Intn(int(d.Sides))
	fmt.Printf("Roll result for %s: %d (random: %d)\n", d, roll, r)
	return
}

func (d *DiceRoll) String() string {
	return fmt.Sprintf("<DiceRoll: %dd%d + %d>", d.Num, d.Sides, d.Modifier)
}

func NewDiceRoll(num uint8, sides uint8, mod uint8) *DiceRoll {
	d := new(DiceRoll)
	d.Num = num
	d.Sides = sides
	d.Modifier = mod
	return d
}
