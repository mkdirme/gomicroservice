package datastore

import (
	"errors"
	"fmt"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//UserStore interafce for all DB interactions
type UserStore interface {
	SaveUser(string, string) (bool, error)
	FindUser(email string) (*interface{}, error)
	// DeleteUser()
	// UpdateUser()
	//SelectAllUsers()
	IsUser(string, string) (bool, error)
}

//User model for  DB
type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	UserName  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
}

//UserMgoStore struct
type UserMgoStore struct {
	C *mgo.Collection
}

//NewUserMgoStore constructor for UserMgoStore class
func NewUserMgoStore() *UserMgoStore {
	store := new(UserMgoStore)
	session, err := mgo.Dial("mongodb://admin:admin1@ds119688.mlab.com:19688/mgo")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	collection := session.DB("mgo").C("users")
	store.C = collection
	return store
}

//ToString Returns String representation of current oject
func (store *UserMgoStore) ToString() string {
	return "" //string(user.FirstName)
}

//FindUser return the User Profile
func (store *UserMgoStore) FindUser(email string) (*interface{}, error) {
	var i interface{}
	return &i, nil //string(user.FirstName)
}

//SaveUser saves curent user as a valid user
func (store *UserMgoStore) SaveUser(email string, password string) (bool, error) {

	usr := struct {
		email    string
		password string
	}{
		email,
		password,
	}
	isUser, err := store.isUser(email)
	if err != nil {
		log.Panicln(err)
		return false, err
	}
	if !isUser {
		err := store.C.Insert(&usr)
		if err != nil {
			return false, err
		}
		return true, nil
	}

	return false, errors.New("this user already exsist. users must be unique")
}

//DeleteUser is used to delete curent user from system
func (store *UserMgoStore) DeleteUser(email string) (bool, error) {
	isUser, err := store.isUser(email)
	if err != nil {
		log.Panicln(err)
		return false, err
	}
	if isUser {
		// Delete record
		err := store.C.Remove(bson.M{"email": email})
		if err != nil {
			fmt.Printf("remove fail %v\n", err)
		}
		return true, nil
	}

	return false, errors.New("user " + email + " dose not exist, and can not be deleted")
}

//IsUser check if user is a valid user
func (store *UserMgoStore) IsUser(email string, password string) (bool, error) {
	usr := struct {
		email    string
		password string
	}{
		email,
		password,
	}
	err := store.C.Find(bson.M{"email": email}).One(&usr)
	if err != nil {
		log.Fatal(err)
	}

	return false, nil
}

func (store *UserMgoStore) isUser(email string) (bool, error) {
	usr := struct {
		email    string
		password string
	}{}

	err := store.C.Find(bson.M{"email": email}).One(&usr)
	if err != nil {
		log.Println(err)
		return false, err
	} else if len(usr.email) <= 0 {
		return false, errors.New("user " + email + " dose not exist")
	}
	return true, nil
}
