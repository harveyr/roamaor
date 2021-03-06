package domain

import (
	"log"
	// "math"
	"math/rand"
)

func ShouldFight(b *Being, tickSeconds float64) (shouldFight bool) {
	shouldFight = false

	// Fights per hour whenn 1.0 danger
	fightsPerHour := float64(2)
	loc := b.SmallestLocation()
	threshold := fightsPerHour / 3600 * tickSeconds * float64(loc.Danger)
	if rand.Float64() < threshold {
		shouldFight = true
	}
	return
}

func TickBeing(b *Being, multiplier float64) {
	// tickSeconds := math.Min(3600, b.SinceLastTick() * multiplier)
	tickSeconds := 60 * multiplier
	log.Printf("--- Ticking %s [%f seconds] ---\n", b, tickSeconds)

	Heal(b, tickSeconds)

	if ShouldFight(b, tickSeconds) {
		mob := NewMob(b.Level)
		EquipBeingChance(mob, 0.3)
		Fight(b, mob)
	}

    Roam(b, tickSeconds)
    ApplyProgress(b)
	b.UpdateLastTick()
}
	
