package fight

import (
	"testing"
	"../models/being"
)

func TestHit(t *testing.T) {
	attacker := being.NewToon("Attacking Toon")
	victim := being.NewToon("Defending Toon")
	initialVictimHp := victim.Hp

	Hit(attacker, victim)

	if victim.Hp == initialVictimHp {
		t.Errorf("Victim's hit points (%d) were not affected.", victim.Hp)
	}
}

func TestFight(t *testing.T) {
	for i := 0; i < 100; i++ {
		attacker := being.NewToon("Attacking Toon")
		victim := being.NewToon("Defending Toon")
		initialVictimHp := victim.Hp
		initialAttackerHp := attacker.Hp

		winner := Fight(attacker, victim)

		if winner == nil {
			t.Errorf("No winner of fight between %s and %s", attacker, victim)
		}

		if (attacker.Hp > initialAttackerHp) {
			t.Errorf("Attacker hp greater than initial: %d > %d", attacker.Hp, initialAttackerHp)
		}
		if (victim.Hp > initialVictimHp) {
			t.Errorf("Victim hp greater than initial: %d > %d", victim.Hp, initialVictimHp)
		}
	}
}
