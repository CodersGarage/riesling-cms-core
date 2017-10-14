package data

import (
	"riesling-cms-core/app/conn"
	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
)

const (
	USER_LEVEL_ADMIN  = 1
	USER_LEVEL_MEMBER = 2
	COLLECTION_NAME   = "users"
)

type User struct {
	bongo.DocumentBase `json:"-",bson:",inline"`
	Hash     string    `json:"hash"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password,omitempty"`
	Level    int       `json:"level"`
}

func (u *User) Save() bool {
	if err := conn.GetConnection().Collection(COLLECTION_NAME).Save(u); err == nil {
		return true
	}
	return false
}

func (u *User) Delete() {

}

func (u *User) LevelUp() {

}

func (u *User) LevelDown() {

}

func (u *User) Ban() {

}

func (u *User) IsEmailExists() bool {
	results := conn.GetConnection().Collection(COLLECTION_NAME).Find(bson.M{
		"email": u.Email,
	})
	if results.Next(u) {
		return true
	}
	return false
}

func (u *User) Count() int {
	count, err := conn.GetConnection().Collection(COLLECTION_NAME).Collection().Count()
	if err != nil {
		return 0
	}
	return count
}

func (u *User) Get(hash string) bool {
	results := conn.GetConnection().Collection(COLLECTION_NAME).Find(bson.M{
		"hash": hash,
	})
	return results.Next(u)
}
