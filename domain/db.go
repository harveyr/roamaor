package domain

import (
	"log"
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

func InsertDoc(collection string, doc DocInterface) bson.ObjectId {
	c := GetCollection(collection)
	id := bson.NewObjectId()
	doc.SetId(id)
	err := c.Insert(doc)
	if err != nil {
		log.Fatal("Failed to insert doc: ", err)
	}
	return id
}

func DeleteDoc(collection string, doc DocInterface) {
	c := GetCollection(collection)
	query := map[string]bson.ObjectId{
		"_id": doc.Id(),
	}
	err := c.Remove(query)
	if err != nil {
		log.Fatalf("Failed to delete doc: %s, %s", collection, doc)
	}
}
