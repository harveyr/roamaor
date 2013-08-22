package domain

import (
    "fmt"
    "log"
    "reflect"
)

type Point struct {
	X, Y float64
}

type Location struct {
	Name string
	Start *Point
    W, H uint16
}

func PointFromMap(m map[string]interface{}) *Point {
	p := Point{}
	if val, ok := m["x"].(float64); ok {
		p.X = val
	} else {
		t := reflect.TypeOf(m["x"])
		log.Fatal("Failed to convert map[x]: ", m["x"], t)
	}
	if val, ok := m["y"].(float64); ok {
		p.Y = val
	} else {
		log.Fatal("Failed to convert map[y]: ", m["y"])
	}
	return &p
}

func NewPoint(x float64, y float64) *Point {
	return &Point{X: x, Y: y}
}

func NewLocation(name string, x uint16, y uint16, w uint16, h uint16) *Location {
	l := &Location{Name: name, W: w, H: h}
	l.Start = NewPoint(float64(x), float64(y))
	return l
}

func (l *Location) String() (repr string) {
    repr = fmt.Sprintf("<[Location] %s>", l.Name)
    return
}

func (l *Location) Center() *Point {
	x := l.Start.X + float64(l.W) / 2
	y := l.Start.Y + float64(l.H) / 2
	return NewPoint(x, y)
}
