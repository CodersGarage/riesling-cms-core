package utils

import "golang.org/x/crypto/bcrypt"

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
