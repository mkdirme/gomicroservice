package datastore

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

var db *sql.DB

//UserMySQLStore struct
type UserMySQLStore struct {
	DB *sql.DB
}

//NewUserMySQLStore constructor for UserMySQLStore class
func NewUserMySQLStore(dataSourceName string) *UserMySQLStore {
	store := new(UserMySQLStore)
	var err error
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Panic(err)
	}
	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	store.DB = db

	return store
}

//ToString Returns String representation of current oject
func (store *UserMySQLStore) ToString() string {
	return "" //string(user.FirstName)
}

//FindUser return the User Profile
func (store *UserMySQLStore) FindUser(email string) (*interface{}, error) {
	var i interface{}
	return &i, nil //string(user.FirstName)
}

//SaveUser saves curent user as a valid user
func (store *UserMySQLStore) SaveUser(email string, password string) (bool, error) {
	var err error
	defer db.Close()
	var isUser bool
	err = db.QueryRow("SELECT IF(COUNT(*),'true','false') FROM users WHERE email = ? AND password = ?", email, password).Scan(&isUser)
	if err != nil {
		fmt.Println(err)
	}
	if isUser {
		return false, errors.New("Invalid Email and Password")
	}
	stmt, err := db.Prepare("INSERT INTO users(fname,lname,email,password) VALUES(?,?,?,?)")
	if err != nil {
		log.Fatal(err)
		return false, errors.New("")
	}
	_, err = stmt.Exec("", "", email, password, "", "")
	if err != nil {
		log.Fatal(err)
		return false, errors.New("")
	}

	return true, nil
}

//DeleteUser is used to delete curent user from system
func (store *UserMySQLStore) DeleteUser(email string) (bool, error) {
	var err error
	defer db.Close()
	var isUser bool
	_, err = db.Query("DELETE FROM User WHERE Email=" + string(email))
	if err != nil {
		log.Println(err)
		return false, errors.New(err.Error())
	}
	if isUser {
		return true, errors.New("Invalid Email")
	}

	return false, nil
}

//IsUser check if user is a valid user
func (store *UserMySQLStore) IsUser(email string) bool {
	var err error
	defer db.Close()
	var isUser bool
	err = db.QueryRow("SELECT IF(COUNT(*),'true','false') FROM user WHERE Email = ?", email).Scan(&isUser)
	//	err = db.QueryRow("SELECT IF(COUNT(*),'true','false') FROM user WHERE Email = ? AND Password = ?", email, password).Scan(&isUser)
	if err != nil {
		log.Println(err)
		return false
	}
	if isUser {
		return true
	}

	return false
}
