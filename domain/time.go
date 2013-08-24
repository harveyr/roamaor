package domain

import (
	"log"
)

func TickBeing(b *Being) {
	log.Printf("--- Ticking %s ---\n", b)
	log.Print("Since last tick: ", b.SinceLastTick())
    Roam(b, b.SinceLastTick())
	b.UpdateLastTick()
}
