// Roamar domain classes 
package domain

import (
	"fmt"
	"math"
	"math/rand"
	"labix.org/v2/mgo/bson"

)

type MongoDocInterface interface {
	SetId(id bson.ObjectId)
	ObjectId() bson.ObjectId
}

type MongoDoc struct {
	Id bson.ObjectId "_id"
}

func (d MongoDoc) SetId(id bson.ObjectId) {
	d.Id = id
}

func (d MongoDoc) ObjectId() bson.ObjectId {
	return d.Id
}

var itemNamePrefixes = []string{
	"Crap",
	"Benign",
	"Slimy",
	"Paltry",
	"Measly",
	"Substandard",
	"Beleaguered",
	"Patchwork",
	"Stinky",
	"Tolerable",
	"Middlin'",
	"Moderate",
	"Acceptable",
	"Intriguing",
	"Notable",
	"Inspiring",
	"Distasteful",
	"Corrupted",
	"Sorrowful",
	"Vengeful",
	"Punishing",
	"Widowing",
	"Sadistic",
}

func FromSliceByLevel(level int, slice []string) string {
	span := 5
	medIndex := math.Min(
		float64(len(slice) - span - 1),
		math.Floor(float64(level) / 100 * float64(len(slice))))
	minIndex := math.Max(0, float64(medIndex - 5))
	maxIndex := math.Min(medIndex + 5, float64(len(slice) - 1))
	index := int(minIndex) + int(math.Floor((maxIndex - minIndex) * rand.Float64()))
	return slice[index]
}

func ItemNamePrefix(level int) string {
	return FromSliceByLevel(level, itemNamePrefixes)
}

func PrefixedItemName(suffix string, level int) string {
	return fmt.Sprintf("%s %s", ItemNamePrefix(level), suffix)
}
