package src

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"strconv"
	"time"
)

func Now() string {
	monthDay, month, hour, min, sec, year := time.Now().Day(), time.Now().Month(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), time.Now().Year()
	date := strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(monthDay)
	clock := strconv.Itoa(hour) + ":" + strconv.Itoa(min) + ":" + strconv.Itoa(sec)
	now := date + " " + clock
	return now
}


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
//
func CreateToken(userid uint64, username string) (string, error) {
	var err error
	//Creating Access Token
	_ = os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["username"] = username
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

//func Authentication(next http.HandlerFunc) http.HandlerFunc { // get invoices and returns in json format
//	return func(w http.ResponseWriter, r *http.Request) {
//		//Login(w http.ResponseWriter, r *http.Request)
//		log.Println("middleware", r.URL)
//		var user = entities.User{
//			ID:       1,
//			Username: "username",
//			Password: "password",
//		}
//		var u entities.User
//		u.Username = r.Header.Get("username")
//		u.Password = r.Header.Get("password")
//		//compare the user from the request, with the one we defined:
//		if user.Username != u.Username || user.Password != u.Password {
//			w.WriteHeader(http.StatusUnauthorized)
//			return
//		}
//		//next(w, r)
//		token, err := CreateToken(user.ID, user.Username)
//		if err != nil {
//			panic(err)
//			return
//		}
//		w.Header().Add("Content-Type", "application/json")
//		w.WriteHeader(http.StatusOK)
//		encoder := json.NewEncoder(w)
//		encoder.SetIndent("", "\t")
//		_ = encoder.Encode(token)
//		next(w, r)
//	}
//}
////