package weapon

import (
	"fmt"
	"../../models"
	"../dice"
)

var names = []string{
	"Stick",
	"Pond Stone",
	"River Stone",
	"Stone",
	"Stone",
	"Pointed Stick",
	"Fig",
	"Tilapia",
	"Stabber",
	"Blarneystone",
}

func RandName(level uint16) string {
	return models.PrefixedName(models.FromSliceByLevel(level, names), level)
}

type Weapon struct {
	Name string
	Level uint16
	Damage *dice.DiceRoll
}

func (w *Weapon) String() string {
	return fmt.Sprintf("<[Weapon] %s [%d]>", w.Name, w.Level)
}

func NewWeapon(level uint16) *Weapon {
	w := new(Weapon)
	w.Name = RandName(level)
	w.Level = level
	w.Damage = dice.NewDiceRoll(2, 6, uint8(level))
	return w
}
