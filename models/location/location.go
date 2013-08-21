package location

import (
    "fmt"
)

type Location struct {
    X, Y, W, H uint16
}


func (l *Location) String() (repr string) {
    repr = fmt.Sprintf("<Location {x: %d, y: %d}>", l.X, l.Y)
    return
}

func (l *Location) SetBounds(x uint16, y uint16, w uint16, h uint16) {
	l.X = x
	l.Y = y
	l.W = w
	l.H = h
}
