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
    X1, Y1, X2, Y2 int
}

func NewLocation(name string, x int, y int, w int, h int) *Location {
	c := GetCollection(LOCATION_COLLECTION)
	l := &Location{Name: name, X1: x, Y1: y, X2: x + w, Y2: y + h}
	l.Id = bson.NewObjectId()

	if err := c.Insert(l); err != nil {
		log.Printf("Failed to insert location %s (%s)", l, err)
	}

	return l
}

func (l *Location) String() (repr string) {
    repr = fmt.Sprintf("<[Location] [%s] {%d, %d}>", l.Name, l.X1, l.Y1)
    return
}

func (l Location) Save() {
	c := GetCollection(LOCATION_COLLECTION)
	if err := c.UpdateId(l.Id, l); err != nil {
		log.Printf("Failed to save location %s (%s)", l, err)
	}
}

func FetchLocationsAt(x float64, y float64) []Location {
	var locs []Location
	query := make(map[string]interface{})
	query["x1"] = map[string]float64{"$lte": x}
	query["x2"] = map[string]float64{"$gte": x}
	query["y1"] = map[string]float64{"$lte": y}
	query["y2"] = map[string]float64{"$gte": y}
	c := GetCollection(LOCATION_COLLECTION)
	c.Find(query).All(&locs)
	return locs
}
