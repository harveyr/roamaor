package domain

import (
	"log"
	"math/rand"
)

func TickBeing(b *Being) {
	log.Printf("--- Ticking %s ---\n", b)
	log.Print("Since last tick: ", b.SinceLastTick())
	tickTime := b.SinceLastTick()

	if rand.Float32() > 0.5 {
		Fight(b, NewMob(b.Level))
	}

    Roam(b, tickTime)
	b.UpdateLastTick()
	log.Print("after b.UpdateLastTick(), last tick: ", b.LastTick)
}
