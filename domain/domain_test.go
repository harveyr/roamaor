package domain

import (
	"testing"
	"reflect"
	"fmt"
	// "labix.org/v2/mgo/bson"
)

const TESTDB = "roamaor_test"

type FakeDoc struct {
	Name string
}

func (f *FakeDoc) String() string {
	return fmt.Sprintf("<FakeDoc: %s>", f.Name)
}

func (f *FakeDoc) Serialize() map[string]interface{} {
	m := make(map[string]interface{})
	m["Name"] = f.Name
	return m
}


func InitTestDb() {
	InitDb("localhost", TESTDB)
}

func TestInitDb(t *testing.T) {
	InitTestDb()
	CloseSession()
}

func TestInsertAndDeleteDoc(t *testing.T) {
	InitTestDb()
	collection := "fakedocs"
	doc := new(FakeDoc)
	doc.Name = "Bongo"
	id := InsertDoc(collection, doc.Serialize())
	if reflect.TypeOf(id).Name() != "ObjectId" {
		panic("Returned id is not an ObjectId")
	}
	DeleteDoc(collection, id)
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
		name := PrefixedItemName(suffix, uint16(i))
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

func TestHit(t *testing.T) {
	attacker := NewToon("Attacking Toon")
	victim := NewToon("Defending Toon")
	initialVictimHp := victim.Hp

	Hit(attacker, victim)

	if victim.Hp == initialVictimHp {
		t.Errorf("Victim's hit points (%d) were not affected.", victim.Hp)
	}
}

func TestFight(t *testing.T) {
	for i := 0; i < 100; i++ {
		attacker := NewToon("Attacking Toon")
		victim := NewToon("Defending Toon")
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


