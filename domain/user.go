package domain

import (
	"fmt"
	"log"
	"strings"
	"labix.org/v2/mgo/bson"
)

const USER_COLLECTION = "users"

type User struct {
	Id bson.ObjectId  "_id"
	ToonId bson.ObjectId `bson:"toonid,omitempty"` 
	Email string
	Name string
}

func CanCreateUser(email string) bool {
	uniqueQuery := make(map[string]interface{})
	uniqueQuery["email"] = strings.ToLower(email)
	return !DocExists(USER_COLLECTION, uniqueQuery)
}

func NewUser(email string) *User {
	if !CanCreateUser(email) {
		log.Print("User already exists for email ", email)
		return nil
	}

	u := User{
		Id:			bson.NewObjectId(),
		Email: strings.ToLower(email),
	}
	c := GetCollection(USER_COLLECTION)
	log.Print("u.Id: ", u.Id)

	err := c.Insert(&u)
	if err != nil {
		log.Printf("[WARNING] Failed to insert user: %s (%s) ", u, err)
	}
	return &u
}

func FetchUser(email string) *User {
	query := map[string]string{
		"email": strings.ToLower(email),
	}
	var user User
	c := GetCollection(USER_COLLECTION)
	err := c.Find(query).One(&user)
	if err != nil {
		log.Printf("Failed to fetch user %s (%s)", email, err)
		return nil
	}
	return &user
}

func (u User) String() string {
	return fmt.Sprintf("<User %s | %s>", u.Id, u.Email)
}

func (u User) Publicize() map[string]interface{} {
	m := make(map[string]interface{})
	m["name"] = u.Name
	m["email"] = u.Email
	m["_id"] = u.Id
	m["toonid"] = u.ToonId
	return m
}

func (u User) Save() {
	c := GetCollection(USER_COLLECTION)
	if err := c.UpdateId(u.Id, u); err != nil {
		log.Printf("Failed to save user %s (%s)", u, err)
	}
}
