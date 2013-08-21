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


func New(name string, x uint16, y uint16, w uint16, h uint16) *Location {
	l := &Location{Name: name, W: w, H: h}
	l.Start = new(Point{X: float32(x), Y: (float32(y))})
	return l
}

func NewPoint(x uint16, y uint16) *Point {
	p := &Point(X: float32(x), y: float32(Y))
	return p
}

func (l *Location) String() (repr string) {
    repr = fmt.Sprintf("<[Location] %s>", l.Name)
    return
}

func (l *Location) Center() *Point {
	return new(Point{
		X: l.Start.X + float32(l.W) / 2,
		Y: l.Start.Y + float32(l.H) / 2,
	})
}

// func (l *Location) Bounds() map[string]float64 {
// 	return map[string]float64{
// 		"X1": float64(l.X),
// 		"Y1": float64(l.Y),
// 		"X2": float64(l.X + l.W),
// 		"Y2": float64(l.Y + l.H),
// 	}
// }
