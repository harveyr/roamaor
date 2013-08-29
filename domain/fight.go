package domain

import (
	"log"
	"fmt"
	"time"
	"math"
	"math/rand"
)

var _ string = fmt.Sprintf("%d", 0)

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
	newHp := float64(b.Hp) + math.Ceil(seconds / 60)
	newHpInt := int(math.Min(float64(b.MaxHp), newHp))
	log.Printf("Hp: %d -> %d", b.Hp, newHp)
	b.Hp = newHpInt
	b.Save()
}

func LogFight(fighter *Being, opponent *Being, victor bool) *LogItem {
	if !fighter.IsToon() {
		// Don't bother logging fights for no-toons
		return nil
	}

	fighter.Fights += 1
	if victor {
		fighter.FightsWon += 1
	}
	fighter.Save()
	
	item := NewLogItem(fighter, LOG_FIGHT)
	item.SetAttr("victor", victor)
	item.SetAttr("opponentName", opponent.Name)
	item.SetAttr("opponentLevel", opponent.Level)
	item.Save()
	return item
}

func WinFight(winner *Being, loser *Being) {
	if winner.IsToon() {
		logItem := LogFight(winner, loser, true)
		if !loser.IsToon() && loser.Weapon.Level > winner.Weapon.Level {
			winner.Weapon = loser.Weapon
			logItem.SetAttr("weaponWonName", winner.Weapon.Name)
			logItem.SetAttr("weaponWonLevel", winner.Weapon.Level)
			logItem.Save()
		}
	}
	if loser.IsToon() {
		LogFight(loser, winner, false)
	}
}

func Fight(attacker *Being, victim *Being) {
	log.Printf("Fight! %s vs %s", attacker, victim)
	rand.Seed(time.Now().UTC().UnixNano())
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
		WinFight(victim, attacker)
	} else if victim.Hp <= 0 {
		WinFight(attacker, victim)
	} else {
		log.Fatal("[Fight] Couldn't determine winner")
	}
	return
}
