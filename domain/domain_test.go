package domain

import (
	"testing"
	// "fmt"
	"log"
	"strings"
	// "labix.org/v2/mgo/bson"
)

const TESTDB = "roamaor_test"

func InitTestDb() {
	InitDb("localhost", TESTDB)
	err := _db.DropDatabase()
	if err != nil {
		log.Fatal("Failed to wipe db")
	}
}

func TestInitDb(t *testing.T) {
	InitTestDb()
	CloseSession()
	InitTestDb()
}

func TestNamePrefixes(t *testing.T) {
	prefix := ItemNamePrefix(1)
	if len(prefix) < 1 {
		t.Errorf("Prefix of 0 length: %s", prefix)
	}
}

func TestPrefixedName(t *testing.T) {
	for i := 0; i < 1000; i++ {
		suffix := "Stabber"
		name := PrefixedItemName(suffix, i)
		if len(name) <= (len(suffix) + 1) {
			t.Errorf("Prefixed name '%s' is no longer than suffix '%s'", name, suffix)
		}
	}
}

func TestWeaponDiceRoll(t *testing.T) {
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

func TestDiceRoll(t *testing.T) {
	attempts := 100
	uniqueCount := 0
	rolls := make([]int, attempts)
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

func TestSaveToonFields(t *testing.T) {
	name := "TestSaveToon"
	nameLower := strings.ToLower(name)
	log.Print("nameLower: ", nameLower)
	toon := NewToon(name)
	SaveFields(BEING_COLLECTION, toon, "Name", "Level")
}

// func TestNewToon(t *testing.T) {
// 	name := "TestNewToon Toon"
// 	toon := NewToon(name)
// 	if toon.Name != name {
// 		log.Fatalf("Toon name does not match input: (%s != %s)", name, toon.Name)
// 	}
// 	if toon.NameLower != strings.ToLower(name) {
// 		log.Fatal("Toon NameLower does not match input")
// 	}

// 	if toon.Level != 1 {
// 		log.Fatalf("Toon.Level is %d (expected 1)", toon.Level)
// 	}
// 	log.Print("toon.Id: ", toon.Id)
// 	DeleteDoc(BEING_COLLECTION, toon)
// }

// func TestFetchAllToons(t *testing.T) {
// 	NewToon("TestFetchAllToons Toon 1")
// 	NewToon("TestFetchAllToons Toon 2")
// 	toons := FetchAllToons()
// 	log.Print("toons: ", toons)
// 	if len(toons) != 2 {
// 		log.Fatal("Expected 2 toons. Fetched ", len(toons))
// 	}
// }

// func TestHit(t *testing.T) {
// 	attacker := NewToon("Attacking Toon")
// 	victim := NewToon("Defending Toon")
// 	initialVictimHp := victim.Hp

// 	Hit(attacker, victim)

// 	if victim.Hp == initialVictimHp {
// 		t.Errorf("Victim's hit points (%d) were not affected.", victim.Hp)
// 	}
// }

// func TestFight(t *testing.T) {
// 	for i := 0; i < 100; i++ {
// 		attacker := NewToon("Attacking Toon")
// 		victim := NewToon("Defending Toon")
// 		initialVictimHp := victim.Hp
// 		initialAttackerHp := attacker.Hp

// 		winner := Fight(attacker, victim)

// 		if winner == nil {
// 			t.Errorf("No winner of fight between %s and %s", attacker, victim)
// 		}

// 		if (attacker.Hp > initialAttackerHp) {
// 			t.Errorf("Attacker hp greater than initial: %d > %d", attacker.Hp, initialAttackerHp)
// 		}
// 		if (victim.Hp > initialVictimHp) {
// 			t.Errorf("Victim hp greater than initial: %d > %d", victim.Hp, initialVictimHp)
// 		}
// 	}
// }


