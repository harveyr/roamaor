package dice

import "testing"

func TestDiceRoll(t *testing.T) {
	for i := 0; i < 100; i++ {
		d := NewDice(2, 6, 2)
		roll := d.Roll()
		if roll < 4 || roll > 14 {
			t.Errorf("Invalid roll result for %s: %d", d, roll)
		}
	}
}
