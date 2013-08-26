package domain

import (
	"fmt"
	"log"
	"time"
	"labix.org/v2/mgo/bson"
)

const (
	LOG_COLLECTION = "toonlogs"
	LOG_FIGHT = iota
)

type LogItem struct {
	Id bson.ObjectId "_id"
	ToonId bson.ObjectId `bson:"toonid,omitempty"` 
	LogType int
	data map[string]interface{}
	Created time.Time
}

func (l *LogItem) String() (s string) {
	s = fmt.Sprintf("<LogItem [Toon %s] [Type %d]>", l.ToonId, l.LogType)
	return
}

func NewLogItem(b *Being, logType int) *LogItem {
	if logType < LOG_FIGHT || logType > LOG_FIGHT {
		panic(fmt.Sprintf("Invalid log type: %d", logType))
	}

	l := new(LogItem)
	l.Id = bson.NewObjectId()
	l.ToonId = b.Id
	l.LogType = logType
	l.Created = time.Now().UTC()

	c := GetCollection(LOG_COLLECTION)
	err := c.Insert(l)
	if err != nil {
		log.Fatalf("[NewLogItem] Failed to insert log item: %s (%s) ", l, err)
	}
	return l
}

func FetchToonLogs(toon *Being) (result []LogItem) {
	query := map[string]bson.ObjectId{"toonid": toon.Id}
	c := GetCollection(LOG_COLLECTION)
	if err := c.Find(query).All(&result); err != nil {
		log.Printf("Failed to fetch logs for toon %s (%s)", toon, err)
	}
	return result
}

func (o LogItem) Save() {
	c := GetCollection(BEING_COLLECTION)
	if err := c.UpdateId(o.Id, o); err != nil {
		log.Printf("Failed to save log item %s (%s)", o, err)
	}
}

