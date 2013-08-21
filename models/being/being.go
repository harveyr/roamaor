package being

import (
    // "log"
    "fmt"
)

type Being struct {
    Name string
    X, Y float32
    Level int
    baseSpeed int
}

func (b *Being) String() (repr string) {
    repr = fmt.Sprintf("<Being: %s {X: %d, Y: %d}>", b.Name, b.X, b.Y)
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

