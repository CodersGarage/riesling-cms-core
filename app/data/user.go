package data

import (
	"riesling-cms-core/app/conn"
	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
)

const (
	USER_LEVEL_ADMIN     = 1
	USER_LEVEL_MEMBER    = 2
	USER_COLLECTION_NAME = "users"
)

type User struct {
	bongo.DocumentBase `json:"-",bson:",inline"`
	Hash     string    `json:"hash"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password,omitempty"`
	Level    int       `json:"level"`
	IsBanned bool      `json:"-"`
}

func (u *User) Save() bool {
	if err := conn.GetConnection().Collection(USER_COLLECTION_NAME).Save(u); err == nil {
		return true
	}
	return false
}

func (u *User) Update(hash string) (bool, *User) {
	preUser := User{}
	if preUser.Get(hash) {
		if len(u.Name) >= 3 && len(u.Name) <= 30 {
			preUser.Name = u.Name
		}
		if len(u.Password) >= 8 && len(u.Password) <= 30 {
			preUser.Password = u.Password
		}
		return preUser.Save(), &preUser
	}
	return false, u
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
	results := conn.GetConnection().Collection(USER_COLLECTION_NAME).Find(bson.M{
		"email": u.Email,
	})
	if results.Next(u) {
		return true
	}
	return false
}

func (u *User) Count() int {
	count, err := conn.GetConnection().Collection(USER_COLLECTION_NAME).Collection().Count()
	if err != nil {
		return 0
	}
	return count
}

func (u *User) Get(hash string) bool {
	results := conn.GetConnection().Collection(USER_COLLECTION_NAME).Find(bson.M{
		"hash": hash,
	})
	return results.Next(u)
}
