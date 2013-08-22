package domain

import (
    "log"
    "fmt"
    "time"
    "strings"
	"labix.org/v2/mgo/bson"
)

const (
	BEING_TOON = iota
	BEING_NPC = iota
	BEING_MOB = iota
)

const BEING_COLLECTION = "beings"

var mobPrefixes = []string{
	"Nice",
	"Cutesy",
	"Benign",
	"Limping",
	"Paltry",
	"Measly",
	"Sticky",
	"Foul",
}

var mobNames = []string{
	"Hen",
	"Kitten",
	"Danish",
	"Cuddlefuzz",
	"Sugarplum",
	"Wildebisht",
}

type Being struct {
	_id bson.ObjectId
    BeingType int
    Name string
    NameLower string
    Level int
    Hp int
    Location *Point
    Destination *Point
    Weapon *Weapon
    BaseSpeed int
    LastTick time.Time
}

func RandMobName(level int) string {
	return fmt.Sprintf(
		"%s %s",
		FromSliceByLevel(level, mobPrefixes),
		FromSliceByLevel(level, mobNames))
}

func BeingFromMap(m map[string]interface{}) *Being {
	b := Being{}
	failures := make([]string, 0)

	if val, ok := m["name"].(string); ok {
		b.Name = val
	} else {
		failures = append(failures, "name")
	}

	if val, ok := m["namelower"].(string); ok {
		b.NameLower = val
	} else {
		failures = append(failures, "namelower")
	}
	
	if val, ok := m["level"].(int); ok {
		b.Level = val
	} else {
		log.Fatalf("BeingFromMap failed to convert level: %s", m["level"])
	}

	locMap, ok := m["location"].(map[string]interface{})
	if !ok {
		log.Fatalf("BeingFromMap failed to convert location to map: %s", m["location"])
	}
	b.Location = PointFromMap(locMap)

	if m["destination"] == nil {
		b.Destination = nil
	} else {
		destMap, ok := m["destination"].(map[string]interface{})
		if !ok {
			log.Fatal("BeingFromMap failed to convert destination to map: ", m["destination"])
		}
		b.Destination = PointFromMap(destMap)
	}
	log.Print("destination: ", m["destination"], m["destination"] == nil)

	if len(failures) > 0 {
		log.Fatalf("Failed to convert map keys to Being values: %s", failures)
	}

	if m["weapon"] == nil {
		b.Weapon = nil
	} else {
		log.Fatal("UNHANDLED! Weapon loading from db")
	}

	return &b
}

func NewToon(name string) *Being {
	nameLower := strings.ToLower(name)
	uniqueQuery := make(map[string]interface{})
	uniqueQuery["namelower"] = nameLower
	uniqueQuery["beingtype"] = BEING_TOON
	// map[string]interface{Name: NameLower, BeingType: BEING_TOON}
	if DocExists(BEING_COLLECTION, uniqueQuery) {
		return nil
	}

	b := Being{
		Name: name,
		NameLower: nameLower,
		BeingType: BEING_TOON,
		Level: 1,
		Hp: 60,
		BaseSpeed: 2,
		Location: NewPoint(0, 0),
	}

	InsertDoc(BEING_COLLECTION, b)
	insertedMap := FetchOne(BEING_COLLECTION, uniqueQuery)
	return BeingFromMap(insertedMap)
}

func NewMob(level int) *Being {
	b := new(Being)
	b.Name = RandMobName(level)
	b.Level = level
	b.Hp = 20 + 10 * level
	return b
}

func (b *Being) String() (repr string) {
    repr = fmt.Sprintf("<[Being] %s>", b.Name)
    return
}

func (b *Being) Speed() (speed float64) {
    speed = float64(b.BaseSpeed) + float64(b.Level)
    return
}

func (b *Being) DamageDice() *DiceRoll {
	if b.Weapon != nil {
		return b.Weapon.Damage
	}
	return NewDiceRoll(1, 6, b.Level)
}

func (b *Being) TakeDamage(damage int) {
	if b.Hp < damage {
		b.Hp = 0
	} else {
		b.Hp -= damage
	}
}

func (b *Being) SetName(name string) {
	b.Name = name
}

func (b *Being) UpdateLastTick() {
	b.LastTick = time.Now()
}

func (b *Being) SinceLastTick() uint16 {
	if b.LastTick.IsZero() {
		return 0
	}
	duration := time.Now().Sub(b.LastTick)
	return uint16(duration/time.Second)
}
