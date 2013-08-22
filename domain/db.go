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

func InsertDoc(collection string, doc map[string]interface{}) bson.ObjectId {
	c := GetCollection(collection)
	id := bson.NewObjectId()
	doc["_id"] = id
	err := c.Insert(doc)
	if err != nil {
		panic(err)
	}
	return id
}


func DeleteDoc(collection string, docId bson.ObjectId) {
	c := GetCollection(collection)
	c.RemoveId(docId)
}
