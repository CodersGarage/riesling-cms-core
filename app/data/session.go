package data

import (
	"github.com/go-bongo/bongo"
	"time"
	"riesling-cms-core/app/conn"
	"gopkg.in/mgo.v2/bson"
)

const (
	SESSION_VALIDITY_TIME   = time.Hour * time.Duration(24)
	SESSION_COLLECTION_NAME = "sessions"
)

type Session struct {
	bongo.DocumentBase     `json:"-",bson:",inline"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpireTime   time.Time `json:"expire_time"`
	Hash         string    `json:"-"`
}

func (s *Session) Save() bool {
	if err := conn.GetConnection().Collection(SESSION_COLLECTION_NAME).Save(s); err != nil {
		return false
	}
	return true
}

func (s *Session) Get(AccessToken string) bool {
	resultSet := conn.GetConnection().Collection(SESSION_COLLECTION_NAME).Find(bson.M{
		"access_token": AccessToken,
	})
	if resultSet.Next(s) {
		return true
	}
	return false
}

func (s *Session) Delete() bool {
	err := conn.GetConnection().Collection(SESSION_COLLECTION_NAME).DeleteDocument(s)
	return err == nil
}
