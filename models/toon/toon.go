package toon

import (
    // "log"
    "fmt"
)

type Toon struct {
    Name string
    X, Y float32
    baseSpeed int
    Level int
}

func (t *Toon) String() (repr string) {
    repr = fmt.Sprintf("<Toon: %s {X: %d, Y: %d}>", t.Name, t.X, t.Y)
    return
}

func (t *Toon) Speed() (speed float64) {
    speed = float64(t.baseSpeed + t.Level)
    return
}

func (t *Toon) SetName(name string) {
	t.Name = name
}

func (t *Toon) SetBaseSpeed(speed int) {
	t.baseSpeed = speed
	return
}

