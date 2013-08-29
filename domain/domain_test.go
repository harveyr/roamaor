package domain

import (
	"testing"
	"time"
	"log"
	"strings"
	"labix.org/v2/mgo/bson"
)

const TESTDB = "roamaor_test"

func DeleteDocument(collection string, id bson.ObjectId) {
	c:= GetCollection(collection)
	if err := c.RemoveId(id); err != nil {
		log.Fatalf("Failed to delete %s %s", collection, id)
	}
}

func InitTestDb() {
	InitDb("localhost", TESTDB)
	err := _db.DropDatabase()
	if err != nil {
		log.Fatalf("Failed to wipe db")
	}
}

func TestInitDb(t *testing.T) {
	InitTestDb()
	CloseSession()
	InitTestDb()
}

func TestNamePrefixes(t *testing.T) {
	prefix := NamePrefix(1)
	if len(prefix) < 1 {
		t.Errorf("Prefix of 0 length: %s", prefix)
	}
}

func TestPrefixedName(t *testing.T) {
	for i := 0; i < 1000; i++ {
		suffix := "Stabber"
		name := PrefixedName(suffix, i)
		if len(name) <= (len(suffix) + 1) {
			t.Errorf("Prefixed name '%s' is no longer than suffix '%s'", name, suffix)
		}
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

func TestCreateAndDeleteToon(t *testing.T) {
	name := "TestNewToon Toon"
	toon := NewToon(name)
	if toon.Name != name {
		t.Errorf("Toon name does not match input: (%s != %s)", name, toon.Name)
	}
	if toon.NameLower != strings.ToLower(name) {
		t.Error("Toon NameLower does not match input")
	}

	if toon.Level != 1 {
		t.Errorf("Toon.Level is %d (expected 1)", toon.Level)
	}

	fetched := FetchToonById(toon.Id)
	if fetched == nil {
		t.Error("Failed fetch 1")
	}
	if fetched.Id != toon.Id {
		t.Error("Id mismatch")
	}
	if fetched.Name != toon.Name {
		t.Error("Name mismatch")
	}
	
	toon.Delete()
	fetched = FetchToonById(toon.Id)
	if fetched != nil {
		t.Error("Fetched deleted toon: ", fetched)
	}
}

func TestFetchLocationsAt(t *testing.T) {
	loc := NewLocation("TestLocation", LOCATION_TOWN, 0.0, 50, 100, 5, 5)
	locs := FetchLocationsAt(20, 20)
	if len(locs) > 0 {
		t.Error("There should be no locations at {20, 20}")
	}

	locs = FetchLocationsAt(55, 105)
	if len(locs) != 1 {
		t.Error("There should be exactly one location at {55, 105}")
	}

	locs = FetchLocationsAt(55, 120)
	if len(locs) != 0 {
		t.Error("There should zero locations at {55, 120}")
	}

	locs = FetchLocationsAt(55, 105)
	if len(locs) != 1 {
		t.Error("There one location at {55, 105}")
	}
	DeleteDocument(LOCATION_COLLECTION, loc.Id)
}

func TestUpdateLocationsVisited(t *testing.T) {
	toon := NewToon("Test Toon")
	locationName := "Test Location"
	loc := NewLocation(locationName, LOCATION_TOWN, 0.0, 100, 200, 10, 10)

	toon.LocX = 10
	toon.LocY = 10

	UpdateLocationsVisited(toon)
	if len(toon.LocationsVisited) > 0 {
		t.Error("Toon should not have visited locations yet")
	}

	toon.LocX = 101
	toon.LocY = 201

	UpdateLocationsVisited(toon)
	if len(toon.LocationsVisited) != 1 {
		t.Error("Toon should have visited the test location")
	}

	logItems := FetchToonLogs(toon)
	if len(logItems) != 1 {
		t.Errorf("Expected one location discovery log item. Found %d", len(logItems))
	}
	item := logItems[0]
	logLocName, ok := item.Data["locationName"].(string)
	if !ok {
		t.Errorf("Failed to convert locationName in log item data: %s", item.Data)
	}
	if logLocName != locationName {
		t.Errorf("logLocName (%s) != locationName %s", logLocName, locationName)
	}

	UpdateLocationsVisited(toon)
	if len(toon.LocationsVisited) != 1 {
		t.Error("Toon has multiples in visited history")
	}

	DeleteDocument(BEING_COLLECTION, toon.Id)
	DeleteDocument(LOCATION_COLLECTION, loc.Id)
}

func TestEarnedLevel(t *testing.T) {
	createdDate := time.Now().UTC().AddDate(0, -1, 0)
	toon := NewToon("Test Toon")
	toon.Created = createdDate
	EarnedLevel(toon)
	DeleteDocument(BEING_COLLECTION, toon.Id)
}

func TestCreateAndFetchLogItem(t *testing.T) {
	toon := NewToon("Test Toon")
	item := NewLogItem(toon, LOG_FIGHT)
	victim := "Barney"

	item.SetAttr("victim", victim)
	item.Save()

	fetched := FetchToonLogs(toon)
	if len(fetched) != 1 {
		t.Errorf("Failed to fetch the log created item: %s", fetched)
	}
	if fetched[0].LogType != LOG_FIGHT {
		t.Errorf("Fetched log item had type %d. Expected %d.", fetched[0].LogType, LOG_FIGHT)
	}
	if fetched[0].Data["victim"] != victim {
		t.Errorf("Fetched log item had victim %s. Expected %s.", fetched[0].Data["victim"], victim)
	}
	DeleteDocument(BEING_COLLECTION, toon.Id)
	DeleteDocument(LOG_COLLECTION, item.Id)
}

func TestHit(t *testing.T) {
	attacker := NewToon("Attacking Toon")
	victim := NewToon("Defending Toon")
	initialVictimHp := victim.Hp

	Hit(attacker, victim)

	if victim.Hp == initialVictimHp {
		t.Errorf("Victim's hit points (%d) were not affected.", victim.Hp)
	}
	DeleteDocument(BEING_COLLECTION, attacker.Id)
	DeleteDocument(BEING_COLLECTION, victim.Id)
}

func TestEquipAndSaveToon(t *testing.T) {
	toon := NewToon("Test Toon")
	EquipBeing(toon)

	if len(toon.Weapon.Name) == 0 {
		t.Error("Weapon has empty name field: ", toon.Weapon)
	}

	if damageRoll := toon.Weapon.Damage.Roll(); damageRoll < 1 {
		t.Errorf("Damage roll was %d. Expected at least 1.", damageRoll)
	}

	toon.Save()

	fetchedToon := FetchToonById(toon.Id)
	if fetchedToon.Weapon.Name != toon.Weapon.Name {
		t.Errorf("Fetched toon's weapon name (%s) != original (%s)", fetchedToon.Weapon.Name, toon.Weapon.Name)
	}
	if fetchedRoll := fetchedToon.Weapon.Damage.Roll(); fetchedRoll < 1 {
		t.Errorf("Fetched toon's weapon damage roll (%s) < 1 ", fetchedRoll)
	}

	DeleteDocument(BEING_COLLECTION, toon.Id)	
}


func TestFightLogUponVictory(t *testing.T) {
	toon := NewToon("Test Toon")
	mob := NewMob(toon.Level)
	toon.Hp = 9999999
	Fight(toon, mob)

	logItems := FetchToonLogs(toon)
	if len(logItems) != 1 {
		t.Errorf("Expected one fight log item. Found %d items", len(logItems))
		return
	}
	item := logItems[0]
	if item.LogType != LOG_FIGHT {
		t.Errorf("Expected fight log item (type %d). Found type %d", LOG_FIGHT, item.LogType)
	}
	logMobName, ok := item.Data["opponentName"].(string)
	if !ok {
		t.Errorf("Unable to retrieve opponentName from log item data: %s", item.Data)
	}
	if logMobName != mob.Name {
		t.Errorf("opponentName (%s) != mob.Name (%s)", logMobName, mob.Name)
	}

	DeleteDocument(BEING_COLLECTION, toon.Id)
	DeleteDocument(LOG_COLLECTION, item.Id)
}

func TestWinnerGetsBetterWeapon(t *testing.T) {
	toon := NewToon("Test Toon")
	toon.Hp = 99999
	mob := NewMob(1)
	EquipBeing(mob)
	mobWeapon := mob.Weapon

	Fight(toon, mob)

	if toon.Weapon != mobWeapon {
		t.Errorf("Toon did not pick up mob's weapon")
	}
	if len(toon.Weapon.Name) < 3 {
		t.Error("Toon does not appear to have a valid weapon: ", toon.Weapon)
	}

	item := FetchToonLogs(toon)[0]
	if weaponWonName, ok := item.Data["weaponWonName"].(string); !ok || weaponWonName != mobWeapon.Name {
		t.Errorf("Log data does not reflect weapon won (%s): %s", mobWeapon, item)		
	}

	DeleteDocument(BEING_COLLECTION, toon.Id)
}
