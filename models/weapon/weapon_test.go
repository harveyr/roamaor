package weapon

import "testing"

func TestDiceRoll(t *testing.T) {
	w := NewWeapon("Test Weapon", 5)
	damageRoll := w.Damage.Roll()

	if damageRoll < 5 || damageRoll > 35 {
		t.Errorf("Invalid damage roll for %s: %d", w, damageRoll)
	}
}
