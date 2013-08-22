package domain

import (
	"fmt"
)

const WEAPON_COLLECTION = "weapons"

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

func RandName(level int) string {
	return PrefixedItemName(FromSliceByLevel(level, names), level)
}

type Weapon struct {
	Name string
	Level int
	Damage *DiceRoll
}

func (w *Weapon) String() string {
	return fmt.Sprintf("<[Weapon] %s [%d]>", w.Name, w.Level)
}

// func FindWeaponNameLevel(name string, level uint16) *Weapon {
// 	c := GetCollection(WEAPON_COLLECTION)
// 	c.Find(bson.M{"name": name}).One(&result)
// }

func NewWeapon(level int) *Weapon {
	name := RandName(level)
	w := new(Weapon)
	w.Name = name
	w.Level = level
	w.Damage = NewDiceRoll(2, 6, level)
	return w
}
