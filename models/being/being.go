package being

import (
    // "log"
    "fmt"
    "../location"
)

const (
	TOON = iota
	NPC = iota
	MOB = iota
)

type Being struct {
    Name string
    Level int
    Location *location.Point
    Destination *location.Point
    baseSpeed int
    beingType uint8
}

func NewToon(name string) *Being {
	t := &Being{Name: name, beingType: TOON}
	t.Location = location.NewPoint(0, 0)
	t.Destination = nil
	return t
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
