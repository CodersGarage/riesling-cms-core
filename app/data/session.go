package data

import (
	"github.com/go-bongo/bongo"
	"time"
	"riesling-cms-core/app/conn"
)

const (
	SESSION_COLLECTION_NAME = "sessions"
)

type Session struct {
	bongo.DocumentBase
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
