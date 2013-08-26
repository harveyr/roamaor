package domain

import (
	"fmt"
	"labix.org/v2/mgo/bson"
)

const (
	LOG_FIGHT = iota
)

type LogItem struct {
	Id bson.ObjectId "_id"
	ToonId bson.ObjectId
	LogType int
	data map[string]interface{}
}

func NewLogItem(logType int) *LogItem {

	if logType < LOG_FIGHT || logType > LOG_FIGHT {
		panic(fmt.Sprintf("Invalid log type: %d", logType))
	}

	l := new(LogItem)
	l.LogType = logType
	l.

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

func (o LogItem) Save() {
	c := GetCollection(BEING_COLLECTION)
	if err := c.UpdateId(o.Id, o); err != nil {
		log.Printf("Failed to save log item %s (%s)", o, err)
	}
}

