package domain

import (
	"log"
	"reflect"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

var _session *mgo.Session = nil
var _db *mgo.Database = nil

func InitDb(host string, dbName string) {
	if _session != nil {
		CloseSession()
	}
	s, err := mgo.Dial(host)
	if err != nil {
		panic(err)
	}
	_session = s
	_session.SetSafe(&mgo.Safe{})
	_db = _session.DB(dbName)
	return
}

func GetCollection(collection string) *mgo.Collection {
	if _db == nil {
		log.Fatal("Db session has not been initialized.")
	}
	return _db.C(collection)
}

func CloseSession() {
	if _session == nil {
		log.Fatal("Db session has not been initialized.")
	}
	_session.Close()
	return
}

func GetQueryObject() *bson.M {
	return new(bson.M)
}

func DocExists(collection string, doc map[string]interface{}) bool {
	// log.Printf("DocExists: [collection] %s [query] %s", collection, doc)
	c := GetCollection(collection)
	count, err := c.Find(doc).Count()
	if err != nil {
		log.Fatalf("Error while checking doc existence: %s", err)
	}
	return count > 0
}

func FetchOne(collection string, query map[string]interface{}) map[string]interface{} {
	c := GetCollection(collection)
	returnMap := make(map[string]interface{})
	err := c.Find(query).One(&returnMap)
	if err != nil {
		log.Fatalf("Error fetching object: [collection] %s [query] %s [err] %s", collection, query, err)
	}
	return returnMap
}


func SaveFields(collection string, doc MongoDocInterface, fields ...string) {
	log.Print("[SaveFields] doc: ", doc)
	log.Print("[SaveFields] fields: ", fields)
	updateMap := make(map[string]interface{})
	updateMap["_id"] = doc.ObjectId()
	structVal := reflect.Indirect(reflect.ValueOf(doc))
	for _, field := range fields {
		updateMap[field] = structVal.FieldByName(field)
	}
	log.Print("updateMap: ", updateMap)
	// if b.Id == nil {
	// 	b.Id := bson.NewObjectId()
	// }
	// c := GetCollection(collection)
	// err := c.Insert(doc)
	// if err != nil {
	// 	log.Fatal("Failed to insert doc: ", err)
	// }
	// return id
}


// func InsertDoc(collection string, doc DocInterface) bson.ObjectId {
// 	c := GetCollection(collection)
// 	id := bson.NewObjectId()
// 	log.Print("[InsertDoc] Setting doc id to new id: ", id)
// 	doc.SetId(id)
// 	err := c.Insert(doc)
// 	if err != nil {
// 		log.Fatal("Failed to insert doc: ", err)
// 	}
// 	return id
// }

// func DeleteDoc(collection string, doc DocInterface) {
// 	if !doc.ObjectId().Valid() {
// 		log.Fatal("[DeleteDoc] Doc has invalid Id: %s", doc.ObjectId())
// 	}

// 	c := GetCollection(collection)
	
// 	emptyQuery := make(map[string]interface{})
// 	var allResults []interface{}
// 	allErr := c.Find(emptyQuery).All(&allResults)
// 	if allErr != nil {
// 		log.Fatal("[DeleteDoc] Coudln't find ANY docs!")
// 	} else {
// 		log.Printf("[DeleteDoc] Found %d docs: %s", len(allResults), allResults)
// 	}

// 	log.Print("string: ", doc.ObjectId().String())
// 	log.Print("hex: ", doc.ObjectId().Hex())
// 	log.Print("machine: ", doc.ObjectId().Machine())

// 	// query := map[string]bson.ObjectId{
// 	// 	"_id": doc.ObjectId(),
// 	// }
// 	query := map[string]bson.ObjectId{
// 		"_id": doc.ObjectId(),
// 	}

// 	count, err := c.Find(query).Count()
// 	if count == 0 {
// 		log.Fatal("[DeleteDoc] Couldn't find doc with query ", query)
// 	}

// 	err = c.Remove(query)
// 	// var findMap map[string]interface{}
// 	// c.FindId
// 	// err := c.RemoveId(doc.Id())
// 	if err != nil {
// 		log.Fatalf("Failed to delete doc: %s, %s, Query: %s. Error: %s", collection, doc, query, err)
// 		// log.Fatalf("Failed to delete doc: %s, %s, Error: %s", collection, doc, err)
// 	}
// }
