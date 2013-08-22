package domain

import (
	"log"
	// "fmt"
	"time"
	"math/rand"
)

func Hit(hitter *Being, hittee *Being) {
	damage := hitter.DamageDice().Roll()
	// fmt.Printf("damage: %d\n", damage)
	hittee.TakeDamage(uint32(damage))
	return
}

func AttackerSwings(round int, attacker *Being, victim *Being) bool {
	if round == 0 {
		return true
	}
	if r := rand.Float32(); r > 0.5 {
		return true;
	}
	return false;
}

func Fight(attacker *Being, victim *Being) *Being {
	rand.Seed(time.Now().UTC().UnixNano())
	round := 0
	for {
		if AttackerSwings(round, attacker, victim) {
			Hit(attacker, victim)
		} else {
			Hit(victim, attacker)
		}
		if attacker.Hp <= 0 {
			return victim
		} else if victim.Hp <= 0 {
			return attacker
		}
		// fmt.Printf("attacker.Hp: %d\n", attacker.Hp)
		// fmt.Printf("victim.Hp: %d\n", victim.Hp)
		round += 1
		if round > 1000 {
			log.Fatal("Too many rounds!")
		}
	}
	return nil
}
