package domain

import (
	"log"
	// "math"
	"math/rand"
)

func ShouldFight(b *Being) (shouldFight bool) {
	shouldFight = false
	if rand.Float32() > 0.5 {
		if float64(b.Hp) / float64(b.MaxHp) >= 0.5 {
			shouldFight = true
		}
	}
	return
}

func TickBeing(b *Being, multiplier float64) {
	// tickTime := math.Min(3600, b.SinceLastTick() * multiplier)
	tickTime := 60 * multiplier
	log.Printf("--- Ticking %s [%f seconds] ---\n", b, tickTime)
	log.Print("b.LocationsVisited: ", b.LocationsVisited)

	Heal(b, tickTime)

	if ShouldFight(b) {
		mob := NewMob(b.Level)
		EquipBeingChance(b, 0.3)
		Fight(b, mob)
	}

    Roam(b, tickTime)
    ApplyProgress(b)
	b.UpdateLastTick()
}
