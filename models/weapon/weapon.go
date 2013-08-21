package weapon

import (
	"../dice"
)

type Weapon struct {
	Name string
	Level uint8
	Damage *dice.DiceRoll
}

func NewWeapon(name string, level int) *Weapon {
	w := new(Weapon)
	w.Name = name
	w.Level = uint8(level)
	w.Damage = dice.NewDiceRoll(2, 6, uint8(level))
	return w
}
