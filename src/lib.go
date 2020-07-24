package src

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)


// HashAndSalt uses GenerateFromPassword to hash & salt pwd.
// MinCost is just an integer constant provided by the bcrypt
// package along with DefaultCost & MaxCost.
// The cost can be any value you want provided it isn't lower
// than the MinCost (4)
// GenerateFromPassword returns a byte slice so we need to
// convert the bytes to a string and return it
func HashAndSalt(pwd string) string {
	pwdByte := []byte(pwd)
	hash, err := bcrypt.GenerateFromPassword(pwdByte, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

// ComparePasswords compares the hash stored in the database
// with the hash generated from the given password
// Since we'll be getting the hashed password from the DB it
// will be a string so we'll need to convert it to a byte slice
func ComparePasswords(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	bytePwd := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePwd)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
