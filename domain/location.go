package domain

import (
    "fmt"
    "log"
	"labix.org/v2/mgo/bson"
)

const (
	LOCATION_COLLECTION = "locations"
	LOCATION_REGION = iota
	LOCATION_MOUNTAIN = iota
	LOCATION_TOWN = iota
)

type Location struct {
	Id        bson.ObjectId "_id"
	Name 	  string
    X1, Y1, X2, Y2 int
    Danger    float32
    LocationType  int
}

func LocationTypes() map[string]int {
	return map[string]int{
		"region": LOCATION_REGION,
		"mountain": LOCATION_MOUNTAIN,
		"town": LOCATION_TOWN,
	}
}

func NewLocation(name string, locType int, danger float32, x int, y int, w int, h int) *Location {
	c := GetCollection(LOCATION_COLLECTION)
	l := &Location{
		Name: name,
		Danger: danger,
		LocationType: locType,
		X1: x,
		Y1: y,
		X2: x + w,
		Y2: y + h,
	}
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

func (l Location) Width() int {
	return l.X2 - l.X1
}

func (l Location) Height() int {
	return l.Y2 - l.Y1
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

func FetchLocationsVisited(b *Being) []Location {
	query := make(map[string]interface{})
	query["_id"] = map[string][]bson.ObjectId{"$in": b.LocationsVisited}
	var result []Location
	c := GetCollection(LOCATION_COLLECTION)
	if err := c.Find(query).All(&result); err != nil {
		log.Print("Failed to fetch locations visited: ", err)
		return make([]Location, 0)
	}
	return result
}
