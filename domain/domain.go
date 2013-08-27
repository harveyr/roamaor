// Roamar domain classes 
package domain

import (
	"log"
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

func (d *MongoDoc) InitId() {
	if d.Id.Valid() {
		log.Fatal("[SetNewId] Doc already has a valid id: ", d.Id)
	}
	d.Id = bson.NewObjectId()
}

var namePrefixes = []string{
	"Crap",
	"Pitiful",
	"Benign",
	"Broken",
	"Slimy",
	"Paltry",
	"Inconsequential",
	"Unremarkable",
	"Measly",
	"Cutesy",
	"Substandard",
	"Beleaguered",
	"Patchwork",
	"Stinky",
	"Sticky",
	"Tolerable",
	"Middlin'",
	"Moderate",
	"Introverted",
	"Irreversed",
	"Cromulent",
	"Silent",
	"Foul",
	"Acceptable",
	"Intriguing",
	"Notable",
	"Lovable",
	"Inspiring",
	"Distasteful",
	"Corrupted",
	"Inverted",
	"Sorrowful",
	"Courageous",
	"Callipygian",
	"Vengeful",
	"Punishing",
	"Sordid",
	"Sardonic",
	"Sacrilicious",
	"Coruscating",
	"Widowing",
	"Embiggened",
	"Splendid",
	"Sumptuous",
	"Sadistic",
	"Resplendent",
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

func NamePrefix(level int) string {
	return FromSliceByLevel(level, namePrefixes)
}

func PrefixedName(suffix string, level int) string {
	return fmt.Sprintf("%s %s", NamePrefix(level), suffix)
}
