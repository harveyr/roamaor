package domain

import (
	"fmt"
	"log"
	"strings"
	"time"
	// "labix.org/v2/mgo/bson"
)

const (
	BEING_TOON = iota
	BEING_NPC  = iota
	BEING_MOB  = iota
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
	MongoDoc
	BeingType int
	Name      string
	NameLower string
	Level     int
	Hp        int
	LocX      float64
	LocY      float64
	DestX     int
	DestY     int
	BaseSpeed int
	LastTick  time.Time
}

// func (b Being) Publicize() map[string]interface{} {
// 	log.Print("Publicizing being: ", b)
// 	m := make(map[string]interface{})
// 	m["id"] = b.Id
// 	m["name"] = b.Name
// 	m["hp"] = b.Hp
// 	m["level"] = b.Level
// 	m["x"] = b.LocX
// 	m["y"] = b.LocY
// 	return m
// }

// func BeingFromMap(m map[string]interface{}) *Being {
// 	b := Being{}
// 	failures := make([]string, 0)

// 	if val, ok := m["_id"].(bson.ObjectId); ok {
// 		b.id = val
// 	} else {
// 		log.Fatalf("BeingFromMap failed to convert ObjectId: %s", m["_id"])
// 	}

// 	if val, ok := m["name"].(string); ok {
// 		b.Name = val
// 	} else {
// 		failures = append(failures, "name")
// 	}

// 	if val, ok := m["namelower"].(string); ok {
// 		b.NameLower = val
// 	} else {
// 		failures = append(failures, "namelower")
// 	}

// 	if val, ok := m["level"].(int); ok {
// 		b.Level = val
// 	} else {
// 		log.Fatalf("BeingFromMap failed to convert level: %s", m["level"])
// 	}

// 	// locMap, ok := m["location"].(map[string]interface{})
// 	// if !ok {
// 	// 	log.Fatalf("BeingFromMap failed to convert location to map: %s", m["location"])
// 	// }
// 	// b.Location = PointFromMap(locMap)

// 	// if m["destination"] == nil {
// 	// 	b.Destination = nil
// 	// } else {
// 	// 	destMap, ok := m["destination"].(map[string]interface{})
// 	// 	if !ok {
// 	// 		log.Fatal("BeingFromMap failed to convert destination to map: ", m["destination"])
// 	// 	}
// 	// 	b.Destination = PointFromMap(destMap)
// 	// }
// 	// log.Print("destination: ", m["destination"], m["destination"] == nil)

// 	if len(failures) > 0 {
// 		log.Fatalf("Failed to convert map keys to Being values: %s", failures)
// 	}

// 	// if m["weapon"] == nil {
// 	// 	b.Weapon = nil
// 	// } else {
// 	// 	log.Fatal("UNHANDLED! Weapon loading from db")
// 	// }

// 	log.Print("BeingFromMap result: ", b)
// 	return &b
// }

func NewToon(name string) *Being {
	nameLower := strings.ToLower(name)
	uniqueQuery := make(map[string]interface{})
	uniqueQuery["namelower"] = nameLower
	uniqueQuery["beingtype"] = BEING_TOON
	if DocExists(BEING_COLLECTION, uniqueQuery) {
		log.Printf("Being with name %s already exists. Returning nil.", name)
		return nil
	}

	b := Being{
		Name:      name,
		NameLower: nameLower,
		BeingType: BEING_TOON,
		Level:     1,
		Hp:        60,
		BaseSpeed: 2,
		LocX:      0,
		LocY:      0,
	}

	// InsertDoc(BEING_COLLECTION, &b)
	return &b
}

func FetchAllToons() []Being {
	query := map[string]int{
		"beingtype": BEING_TOON,
	}
	var result []Being
	c := GetCollection(BEING_COLLECTION)
	err := c.Find(query).All(&result)
	if err != nil {
		log.Print("WARNING: Failed to fetch all toons")
	}
	return result
}

func RandMobName(level int) string {
	return fmt.Sprintf(
		"%s %s",
		FromSliceByLevel(level, mobPrefixes),
		FromSliceByLevel(level, mobNames))
}

func NewMob(level int) *Being {
	b := new(Being)
	b.Name = RandMobName(level)
	b.Level = level
	b.Hp = 20 + 10*level
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

func (b Being) SinceLastTick() uint16 {
	if b.LastTick.IsZero() {
		return 0
	}
	duration := time.Now().Sub(b.LastTick)
	return uint16(duration / time.Second)
}

