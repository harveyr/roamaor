package domain

import (
	"log"
	"math/rand"
)

func TickBeing(b *Being) {
	tickTime := b.SinceLastTick()
	log.Printf("--- Ticking %s [%f seconds] ---\n", b, tickTime)
	log.Print("b.LocationsVisited: ", b.LocationsVisited)
	if rand.Float32() > 0.5 {
		Fight(b, NewMob(b.Level))
	}

    Roam(b, tickTime)
	b.UpdateLastTick()
}
