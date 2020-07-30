package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"nf_stn/config"
	"os"
	"strconv"
	"strings"

	//"fmt"
	//"github.com/dgrijalva/jwt-go"
	"net/http"
	"nf_stn/entities"
	"nf_stn/src"
	//"os"
	//"strconv"
	//"strings"
)

//func Authentication(next http.HandlerFunc) http.HandlerFunc { // get invoices and returns in json format
//	return func(w http.ResponseWriter, r *http.Request) {
//		var user = entities.User{
//			ID:       1,
//			Username: "username",
//			Password: "password",
//		}
//		var u entities.User
//		u.Username = r.Header.Get("username")
//		u.Password = r.Header.Get("password")
//		//bearToken := r.Header.Get("Authorization")
//		////normally Authorization the_token_xxx
//		//strArr := strings.Split(bearToken, " ")
//		//fmt.Println(strArr)
//		//
//		//token, _ := jwt.Parse(bearToken, func(token *jwt.Token) (interface{}, error) {
//		//	//Make sure that the token method conform to "SigningMethodHMAC"
//		//	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//		//		//return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
//		//	}
//		//	return []byte(os.Getenv("ACCESS_SECRET")), nil
//		//})
//		//if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
//		//	//return
//		//}
//		////token, err := VerifyToken(r)
//		////if err != nil {
//		////	return nil, err
//		////}
//		//var acc AccessDetails
//		//claims, ok := token.Claims.(jwt.MapClaims)
//		//if ok && token.Valid {
//		//	accessUuid, ok := claims["access_uuid"].(string)
//		//	if !ok {
//		//		//return nil, err
//		//	}
//		//	userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
//		//	if err != nil {
//		//		//return nil, err
//		//	}
//		//	acc.AccessUuid = accessUuid
//		//	acc.UserId = userId
//		//	}
//
//
//		//if err != nil {
//		//	return nil, err
//		//}
//		//if len(strArr) == 2 {
//		//	return strArr[1]
//		//}
//
//
//		//compare the user from the request, with the one we defined:
//		if user.Username != u.Username || user.Password != u.Password {
//			w.WriteHeader(http.StatusUnauthorized)
//			return
//		}
//		ts, err := src.CreateToken(user.ID, u.Username)
//		if err != nil {
//			panic(err)
//			return
//		}
//
//		saveErr := src.CreateAuth(user.ID, ts)
//		if saveErr != nil {
//			w.WriteHeader(http.StatusUnprocessableEntity)
//		}
//		tokens := map[string]string{
//			"access_token":  ts.AccessToken,
//			"refresh_token": ts.RefreshToken,
//		}
//
//		w.Header().Add("Content-Type", "application/json")
//		w.WriteHeader(http.StatusOK)
//		encoder := json.NewEncoder(w)
//		encoder.SetIndent("", "\t")
//		_ = encoder.Encode(tokens)
//		next(w, r)
//	}
//}
////

//
func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
//
func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
//
func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}
//
func ExtractTokenMetadata(r *http.Request) (*entities.AccessDetails, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &entities.AccessDetails{
			AccessUuid: accessUuid,
			UserId:   userId,
		}, nil
	}
	return nil, err
}
//
func FetchAuth(authD *entities.AccessDetails) (uint64, error) {
	userid, err := config.Client.Get(authD.AccessUuid).Result()
	if err != nil {
		return 0, err
	}
	userID, _ := strconv.ParseUint(userid, 10, 64)
	return userID, nil
}
//
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var td *entities.Todo
	r.Header.Set("Content-Type", "application/json")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	// Unmarshal
	err = json.Unmarshal(b, &td)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	tokenAuth, err := ExtractTokenMetadata(r)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "\t")
		_ = encoder.Encode("Unauthorized")
		return
	}
	userId, err := FetchAuth(tokenAuth)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "\t")
		_ = encoder.Encode("Unauthorized")
		return
	} else {
		td.UserID = userId

		//you can proceed to save the Todo to a database
		//but we will just return it to the caller here:
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "\t")
		_ = encoder.Encode(td)
	}
}
//
func Logout(w http.ResponseWriter, r *http.Request) {
	tokenAuth, err := ExtractTokenMetadata(r)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "\t")
		_ = encoder.Encode("unauthorized")
		return
	}
	deleted, delErr := src.DeleteAuth(tokenAuth.AccessUuid)
	if delErr != nil || deleted == 0 { //if any goes wrong
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "\t")
		_ = encoder.Encode("unauthorized")
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	_ = encoder.Encode("Successfully logged out")
}
//
func TokenAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := TokenValid(r)
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			info := map[string]string{
				"authentication status":  "unauthorized",
				"method": r.Method,
				"content-type": "application/json",
			}
			encoder := json.NewEncoder(w)
			encoder.SetIndent("", "\t")
			_ = encoder.Encode(info)
			return
		} else {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			info := map[string]string{
				"authentication status": "authorized",
				"method":                r.Method,
				"content-type":          "application/json",
			}
			encoder := json.NewEncoder(w)
			encoder.SetIndent("", "\t")
			_ = encoder.Encode(info)
			next(w, r)
		}
	}
}
//