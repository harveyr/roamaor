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
	"Limping",
}

var mobNames = []string{
	"Hen",
	"Kitten",
	"Rhododendron",
	"Danish",
	"Spore",
	"Cranapple",
	"Self-Reflection",
	"Cuddlefuzz",
	"Sugarplum",
	"Neckbeard",
	"Tilapia",
	"Wildebisht",
	"Sea Cap'n",
	"Bunyip",
}

type Being struct {
	Id        bson.ObjectId "_id"
	BeingType int
	LastTick  time.Time
	Created   time.Time
	Name      string
	NameLower string
	Level     int
	MaxHp     int
	Hp        int
	LocX      float64
	LocY      float64
	DestX     float64
	DestY     float64
	BaseSpeed int
	Fights    int
	FightsWon int
	LocationsVisited []bson.ObjectId "omitempty"
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

	b := new(Being)
	b.Id = bson.NewObjectId()
	b.Level = 1
	b.Name = name
	b.NameLower = strings.ToLower(name)
	b.BeingType = BEING_TOON
	b.MaxHp = 60
	b.Hp = 60
	b.BaseSpeed = 2
	b.Created = time.Now().UTC()

	c := GetCollection(BEING_COLLECTION)
	err := c.Insert(b)
	if err != nil {
		log.Fatalf("[NewToon] Failed to insert toon: %s (%s) ", b, err)
	}
	return b
}

func FetchToonById(id bson.ObjectId) (*Being) {
	c := GetCollection(BEING_COLLECTION)
	b := Being{}
	err := c.FindId(id).One(&b)
	if err != nil {
		log.Printf("Failed to fetch toon with id %s (%s)", id, err)
		return nil
	}
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
		NamePrefix(level),
		FromSliceByLevel(level, mobNames))
}

func NewMob(level int) *Being {
	b := new(Being)
	b.Name = RandMobName(level)
	b.Level = level
	b.Hp = 20 + 10 * level
	b.BeingType = BEING_MOB
	b.Created = time.Now().UTC()
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
	b.Save()
}

func (b *Being) SinceLastTick() float64 {
	if b.LastTick.IsZero() {
		return 0
	}
	duration := time.Now().Sub(b.LastTick)
	return float64(duration) / float64(time.Second)
}

func (b Being) Delete() {
	c := GetCollection(BEING_COLLECTION)
	err := c.RemoveId(b.Id)
	if err != nil {
		log.Fatal("Failed to delete being %s (%s)", b, err)
	}
}

func (b Being) Save() {
	c := GetCollection(BEING_COLLECTION)
	if err := c.UpdateId(b.Id, b); err != nil {
		log.Printf("Failed to save being %s (%s)", b, err)
	}
}

func (b *Being) Reload() {
	c := GetCollection(BEING_COLLECTION)
	c.FindId(b.Id).One(b)
}

func (b Being) IsToon() bool {
	return b.BeingType == BEING_TOON
}
