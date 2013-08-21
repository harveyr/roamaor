package being

import (
    // "log"
    "fmt"
    "../location"
    "../weapon"
    "../dice"
)

const (
	TOON = iota
	NPC = iota
	MOB = iota
)

type Being struct {
    Name string
    Level uint16
    Hp uint32
    Location *location.Point
    Destination *location.Point
    Weapon *weapon.Weapon
    baseSpeed uint8
    beingType uint8
}

func NewToon(name string) *Being {
	b := &Being{Name: name, beingType: TOON}
	b.Location = location.NewPoint(0, 0)
	b.Destination = nil
	b.Weapon = nil
	b.Level = 1
	b.Hp = 40 + 20 * uint32(b.Level)
	return b
}

func (b *Being) String() (repr string) {
    repr = fmt.Sprintf("<[Being] %s>", b.Name)
    return
}

func (b *Being) Speed() (speed float64) {
    speed = float64(b.baseSpeed) + float64(b.Level)
    return
}

func (b *Being) DamageDice() *dice.DiceRoll {
	if b.Weapon != nil {
		return b.Weapon.Damage
	}
	return dice.NewDiceRoll(1, 6, uint8(b.Level))
}

func (b *Being) TakeDamage(damage uint32) {
	if b.Hp < damage {
		b.Hp = 0
	} else {
		b.Hp -= damage
	}
}

func (b *Being) SetName(name string) {
	b.Name = name
}

func (b *Being) SetBaseSpeed(speed uint8) {
	b.baseSpeed = speed
	return
}
