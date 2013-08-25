package domain

import (
    "fmt"
    "log"
	"labix.org/v2/mgo/bson"
)

const (
	WORLD_WIDTH = 600
	WORLD_HEIGHT  = 600
	LOCATION_COLLECTION = "locations"
)

type Location struct {
	Id        bson.ObjectId "_id"
	Name string
    CX, CY, W, H int
}

func NewLocation(name string, x int, y int, w int, h int) *Location {
	c := GetCollection(LOCATION_COLLECTION)
	l := &Location{Name: name, CX: x, CY: y, W: w, H: h}
	l.Id = bson.NewObjectId()

	if err := c.Insert(l); err != nil {
		log.Printf("Failed to insert location %s (%s)", l, err)
	}

	return l
}

func (l *Location) String() (repr string) {
    repr = fmt.Sprintf("<[Location] [%s] {%d, %d}>", l.Name, l.CX, l.CY)
    return
}

func (l Location) Save() {
	c := GetCollection(LOCATION_COLLECTION)
	if err := c.UpdateId(l.Id, l); err != nil {
		log.Printf("Failed to save location %s (%s)", l, err)
	}
}

