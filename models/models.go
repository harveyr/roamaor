package models

import (
	"rand"
)

type Bounded interface {
	Bounds() map[string]float64
}

