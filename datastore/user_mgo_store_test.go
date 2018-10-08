package datastore

import (
	"fmt"
	"log"
	"testing"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Email    string
	Password string
}

func TestDB(t *testing.T) {
	session, err := mgo.Dial("mongodb://admin:admin1@ds119688.mlab.com:19688/mgo")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("mgo").C("users")
	err = c.Insert(&User{"temp@gmail.com", "555381169639"})
	if err != nil {
		log.Fatal(err)
	}

	result := User{}
	err = c.Find(bson.M{"email": "temp@gmail.com"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Email:", result.Email)
}
