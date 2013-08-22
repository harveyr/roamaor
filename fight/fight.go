package fight

import (
	"log"
	"fmt"
	"time"
	"math/rand"
	"../models/being"
)

func Hit(hitter *being.Being, hittee *being.Being) {
	damage := hitter.DamageDice().Roll()
	// fmt.Printf("damage: %d\n", damage)
	hittee.TakeDamage(uint32(damage))
	return
}

func AttackerSwings(round int, attacker *being.Being, victim *being.Being) bool {
	if round == 0 {
		return true
	}
	if r := rand.Float32(); r > 0.5 {
		return true;
	}
	return false;
}

func Fight(attacker *being.Being, victim *being.Being) *being.Being {
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
		fmt.Printf("attacker.Hp: %d\n", attacker.Hp)
		fmt.Printf("victim.Hp: %d\n", victim.Hp)
		round += 1
		if round > 1000 {
			log.Fatal("Too many rounds!")
		}
	}
	return nil
}
