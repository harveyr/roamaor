package location

import (
    "fmt"
)

type Point struct {
	X, Y float32
}

type Location struct {
	Name string
	Start *Point
    W, H uint16
}


func NewPoint(x float32, y float32) *Point {
	return &Point{X: x, Y: y}
}

func New(name string, x uint16, y uint16, w uint16, h uint16) *Location {
	l := &Location{Name: name, W: w, H: h}
	l.Start = NewPoint(float32(x), float32(y))
	return l
}

func (l *Location) String() (repr string) {
    repr = fmt.Sprintf("<[Location] %s>", l.Name)
    return
}

func (l *Location) Center() *Point {
	x := l.Start.X + float32(l.W) / 2
	y := l.Start.Y + float32(l.H) / 2
	return NewPoint(x, y)
}
