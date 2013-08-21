package dice

import "testing"

func TestDiceRoll(t *testing.T) {
	attempts := 100
	uniqueCount := 0
	rolls := make([]uint16, attempts)
	for i := 0; i < attempts; i++ {
		d := NewDiceRoll(2, 6, 2)
		roll := d.Roll()
		if roll < 4 || roll > 14 {
			t.Errorf("Invalid roll result for %s: %d", d, roll)
		}

		unique := true
		for _, r := range rolls {
			if r == roll {
				unique = false
			}
		}
		if unique {
			uniqueCount += 1
		}
		rolls[i] = roll
	}

	if uniqueCount < 6 {
		t.Errorf("Only %d unique rolls were thrown.", uniqueCount)
	}

}
