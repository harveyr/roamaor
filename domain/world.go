package domain

import "log"

const (
	WORLD_WIDTH = 2000
	WORLD_HEIGHT  = 2000
)

func CreateOrUpdateLocation(name string, locType int, danger float32, x int, y int, w int, h int) {
	c := GetCollection(LOCATION_COLLECTION)
	var loc Location
	err := c.Find(map[string]string{"name": name}).One(&loc)
	if err == nil {
		log.Print("Updating location ", name)
		loc.X1 = x
		loc.Y1 = y
		loc.X2 = x + w
		loc.Y2 = y + h
		loc.Save()
	} else {
		log.Printf("Creating location %s (%s)", name, err)
		NewLocation(name, locType, danger, x, y, w, h)
	}
}

func InitWorldLocations() {
	print("\n")
	log.Print("--- Initializing World Locations ---")


	CreateOrUpdateLocation("Lowlunds", LOCATION_REGION, 0.2, 0, 0, 200, 200)
	CreateOrUpdateLocation("Newbville", LOCATION_TOWN, 0.0, 20, 20, 20, 20)
	CreateOrUpdateLocation("Tonky", LOCATION_TOWN, 0.5, 100, 100, 50, 50)

	log.Print("------------------------------------")
	print("\n")
}
