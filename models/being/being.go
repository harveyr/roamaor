package being

import (
    // "log"
    "fmt"
    "time"
    "../location"
    "../weapon"
    "../dice"
    "../../models"
)

const (
	TOON = iota
	NPC = iota
	MOB = iota
)

var mobPrefixes = []string{
	"Nice",
	"Cutesy",
	"Benign",
	"Limping",
	"Paltry",
	"Measly",
	"Sticky",
	"Foul",
}

var mobNames = []string{
	"Hen",
	"Kitten",
	"Danish",
	"Cuddlefuzz",
	"Sugarplum",
	"Wildebisht",
}

type Being struct {
    Name string
    Level uint16
    Hp uint32
    Location *location.Point
    Destination *location.Point
    Weapon *weapon.Weapon
    baseSpeed uint8
    beingType uint8
    LastTick time.Time
}

func RandMobName(level uint16) string {
	return fmt.Sprintf(
		"%s %s",
		models.FromSliceByLevel(level, mobPrefixes),
		models.FromSliceByLevel(level, mobNames))
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

func NewMob(level uint16) *Being {
	b := new(Being)
	b.Name = RandMobName(level)
	b.Level = level
	b.Hp = 20 + 10 * uint32(level)
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

func (b *Being) UpdateLastTick() {
	b.LastTick = time.Now()
}

func (b *Being) SinceLastTick() uint16 {
	if b.LastTick.IsZero() {
		return 0
	}
	duration := time.Now().Sub(b.LastTick)
	return uint16(duration/time.Second)
}
