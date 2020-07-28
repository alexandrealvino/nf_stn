package middleware

import (
	"encoding/json"
	//"fmt"
	//"github.com/dgrijalva/jwt-go"
	"net/http"
	"nf_stn/entities"
	"nf_stn/src"
	//"os"
	//"strconv"
	//"strings"
)

type AccessDetails struct {
	AccessUuid string
	UserId   uint64
}

func Authentication(next http.HandlerFunc) http.HandlerFunc { // get invoices and returns in json format
	return func(w http.ResponseWriter, r *http.Request) {
		var user = entities.User{
			ID:       1,
			Username: "username",
			Password: "password",
		}
		var u entities.User
		u.Username = r.Header.Get("username")
		u.Password = r.Header.Get("password")
		//bearToken := r.Header.Get("Authorization")
		////normally Authorization the_token_xxx
		//strArr := strings.Split(bearToken, " ")
		//fmt.Println(strArr)
		//
		//token, _ := jwt.Parse(bearToken, func(token *jwt.Token) (interface{}, error) {
		//	//Make sure that the token method conform to "SigningMethodHMAC"
		//	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		//		//return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		//	}
		//	return []byte(os.Getenv("ACCESS_SECRET")), nil
		//})
		//if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		//	//return
		//}
		////token, err := VerifyToken(r)
		////if err != nil {
		////	return nil, err
		////}
		//var acc AccessDetails
		//claims, ok := token.Claims.(jwt.MapClaims)
		//if ok && token.Valid {
		//	accessUuid, ok := claims["access_uuid"].(string)
		//	if !ok {
		//		//return nil, err
		//	}
		//	userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		//	if err != nil {
		//		//return nil, err
		//	}
		//	acc.AccessUuid = accessUuid
		//	acc.UserId = userId
		//	}


		//if err != nil {
		//	return nil, err
		//}
		//if len(strArr) == 2 {
		//	return strArr[1]
		//}


		//compare the user from the request, with the one we defined:
		if user.Username != u.Username || user.Password != u.Password {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		token, err := src.CreateToken(user.ID, u.Username)
		if err != nil {
			panic(err)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "\t")
		_ = encoder.Encode(token)
		next(w, r)
	}
}
//