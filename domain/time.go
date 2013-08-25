package domain

import (
	"log"
	// "math"
	"math/rand"
)

func TickBeing(b *Being, multiplier float64) {
	// tickTime := math.Min(3600, b.SinceLastTick() * multiplier)
	tickTime := 60 * multiplier
	log.Printf("--- Ticking %s [%f seconds] ---\n", b, tickTime)
	log.Print("b.LocationsVisited: ", b.LocationsVisited)
	if rand.Float32() > 0.5 {
		Fight(b, NewMob(b.Level))
	}
    Roam(b, tickTime)
	b.UpdateLastTick()
}
