package domain

import (
	// "log"
	"fmt"
)

const WEAPON_COLLECTION = "weapons"

var weaponNames = []string{
	"Stick",
	"Pondstone",
	"Riverstone",
	"Caress",
	"Rag",
	"Stone",
	"Cake Pan",
	"Gelatin",
	"Sweatpalm",
	"Fern",
	"Bucket",
	"Teacup",
	"Breadstick",
	"Pointed Stick",
	"Fig",
	"Tilapia",
	"Dagger",
	"Stabber",
	"Blarneystone",
	"Stinksword",
	"Shartsword",
	"Shortsword",
	"Sword",
	"Longsword",
	"Broadsword",
	"Broadsword",
}

func WeaponName(level int) string {
	return PrefixedName(FromSliceByLevel(level, weaponNames), level)
}

type Weapon struct {
	Name  string
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

func LevelWeapon(level int) (w Weapon) {
	return Weapon{
		Name: WeaponName(level),
		Level: level,
		Damage: LevelDiceRoll(level),
	}
}


func EquipBeing(b *Being) {
	weapon := LevelWeapon(b.Level)
	b.Weapon = weapon
}
