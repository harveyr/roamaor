package weapon

import "testing"

func TestDiceRoll(t *testing.T) {
	w := NewWeapon(5)
	damageRoll := w.Damage.Roll()

	if damageRoll < 5 || damageRoll > 35 {
		t.Errorf("Invalid damage roll for %s: %d", w, damageRoll)
	}
}


func TestRandName(t *testing.T) {
	name := RandName(1)
	if len(name) < 3 {
		t.Errorf("Invalid random name: %s", name)
	}
}
