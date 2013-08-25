package domain

func InitWorldLocations() {
	c := GetCollection(LOCATION_COLLECTION)
	c.DropCollection()
	NewLocation("Newbville", 20, 20, 5, 5)
}
