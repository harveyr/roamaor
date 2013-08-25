package domain

import "log"

func CreateOrUpdateLocation(name string, x int, y int, w int, h int) {
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
		NewLocation(name, x, y, w, h)
	}
}

func InitWorldLocations() {
	print("\n")
	log.Print("--- Initializing World Locations ---")

	CreateOrUpdateLocation("Newbville", 20, 20, 5, 5)
	CreateOrUpdateLocation("Tonky", 30, 30, 5, 5)

	log.Print("------------------------------------")
	print("\n")
}
