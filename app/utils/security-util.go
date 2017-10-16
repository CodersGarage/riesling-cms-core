package utils

import (
	"golang.org/x/crypto/bcrypt"
	"fmt"
	"crypto/sha256"
)

func HashPassword(plain string) string {
	securePass, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(securePass)
}

func CompareHashedPassword(securePass string, plain string) bool {
	return bcrypt.CompareHashAndPassword([]byte(securePass), []byte(plain)) == nil
}

func GetSHA256Hashed(plain string) string {
	hash := sha256.New()
	hash.Write([]byte(plain))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
