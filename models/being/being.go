package being

import (
    // "log"
    "fmt"
)

const (
	TOON = iota
	NPC = iota
	MOB = iota
)

type Being struct {
    Name string
    X, Y float32
    Level int
    baseSpeed int
    beingType uint8
}

func NewToon(name string) *Being {
	return &Being{Name: name, beingType: TOON}
}

func (b *Being) String() (repr string) {
    repr = fmt.Sprintf("<[Being] %s>", b.Name)
    return
}

func (b *Being) Speed() (speed float64) {
    speed = float64(b.baseSpeed + b.Level)
    return
}

func (b *Being) SetName(name string) {
	b.Name = name
}

func (b *Being) SetBaseSpeed(speed int) {
	b.baseSpeed = speed
	return
}

func (b *Being) Bounds() map[string]float64 {
	return map[string]float64{
		"X1": float64(b.X),
		"Y1": float64(b.Y),
		"X2": float64(b.X),
		"Y2": float64(b.Y),
	}
}

