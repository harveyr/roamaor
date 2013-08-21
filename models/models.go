package models

import (
	"fmt"
	"math"
	"math/rand"
)

type Bounded interface {
	Bounds() map[string]float64
}

var namePrefixes = []string{
	"Crap",
	"Benign",
	"Slimy",
	"Substandard",
	"Beleaguered",
	"Patchwork",
	"Stinky",
	"Tolerable",
	"Middlin'",
	"Moderate",
	"Acceptable",
	"Intriguing",
	"Notable",
	"Inspiring",
	"Distasteful",
	"Corrupted",
	"Sorrowful",
	"Vengeful",
	"Punishing",
	"Widowing",
}

func FromSliceByLevel(level uint16, slice []string) string {
	span := 5
	medIndex := math.Min(
		float64(len(namePrefixes) - span - 1),
		math.Floor(float64(level) / 100 * float64(len(slice))))
	minIndex := math.Max(0, float64(medIndex - 5))
	maxIndex := math.Min(medIndex + 5, float64(len(slice) - 1))
	index := int(minIndex) + int(math.Floor((maxIndex - minIndex) * rand.Float64()))
	return slice[index]
}

func NamePrefix(level uint16) string {
	return FromSliceByLevel(level, namePrefixes)
}

func PrefixedName(suffix string, level uint16) string {
	return fmt.Sprintf("%s %s", NamePrefix(level), suffix)
}
