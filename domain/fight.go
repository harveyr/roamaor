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
	hittee.TakeDamage(damage)
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

func Heal(b *Being, seconds float64) {
	newHp := b.Hp + seconds / 60
}

func Fight(attacker *Being, victim *Being) {
	log.Printf("Fight! %s vs %s", attacker, victim)
	rand.Seed(time.Now().UTC().UnixNano())
	attacker.Fights += 1
	victim.Fights += 1
	round := 0
	for {
		if AttackerSwings(round, attacker, victim) {
			Hit(attacker, victim)
		} else {
			Hit(victim, attacker)
		}
		if attacker.Hp <= 0 || victim.Hp <= 0 {
			break
		}
		round += 1
		if round > 1000 {
			log.Fatal("Too many rounds!")
		}
	}

	if attacker.Hp <= 0 {
		victim.FightsWon += 1
	} else if victim.Hp <= 0 {
		attacker.FightsWon += 1
	} else {
		log.Fatal("[Fight] Couldn't determine winner")
	}
	if attacker.IsToon() {
		attacker.Save()
	}
	if victim.IsToon() {
		victim.Save()
	}
	return
}
