package domain

import (
	"fmt"
	"log"
	"strings"
	"time"
	"labix.org/v2/mgo/bson"
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
	Id        bson.ObjectId "_id"
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

func CanCreateToon(name string) bool {
	nameLower := strings.ToLower(name)
	uniqueQuery := make(map[string]interface{})
	uniqueQuery["namelower"] = nameLower
	uniqueQuery["beingtype"] = BEING_TOON
	return !DocExists(BEING_COLLECTION, uniqueQuery)
}

func NewToon(name string) *Being {
	if !CanCreateToon(name) {
		log.Fatal("Can't create toon! Did you check?")
	}

	b := Being{
		Id:			bson.NewObjectId(),
		Name:      name,
		NameLower: strings.ToLower(name),
		BeingType: BEING_TOON,
		Level:     1,
		Hp:        60,
		BaseSpeed: 2,
		LocX:      0,
		LocY:      0,
	}
	log.Print("b.Id: ", b.Id)
	c := GetCollection(BEING_COLLECTION)
	err := c.Insert(b)
	if err != nil {
		log.Fatalf("[NewToon] Failed to insert toon: %s (%s) ", b, err)
	}
	return &b
}

func FetchToonById(id bson.ObjectId) (*Being, error) {
	c := GetCollection(BEING_COLLECTION)
	b := Being{}
	err := c.FindId(id).One(&b)
	if err != nil {
		log.Print("Failed to fetch toon with id ", id)
	}
	return &b, err
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

func (b Being) Delete() {
	c := GetCollection(BEING_COLLECTION)
	err := c.RemoveId(b.Id)
	if err != nil {
		log.Fatal("Failed to delete being %s (%s)", b, err)
	}
}
