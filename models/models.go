package models

type Bounded interface {
	Bounds() map[string]float64
}
