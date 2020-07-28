package middleware

import (
	"encoding/json"
	"net/http"
	"nf_stn/entities"
	"nf_stn/src"
)

func Authentication(next http.HandlerFunc) http.HandlerFunc { // get invoices and returns in json format
	return func(w http.ResponseWriter, r *http.Request) {
		//Login(w http.ResponseWriter, r *http.Request)
		//log.Println("middleware", r.URL)
		var user = entities.User{
			ID:       1,
			Username: "username",
			Password: "password",
		}
		var u entities.User
		u.Username = r.Header.Get("username")
		u.Password = r.Header.Get("password")
		//compare the user from the request, with the one we defined:
		if user.Username != u.Username || user.Password != u.Password {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		//next(w, r)
		token, err := src.CreateToken(user.ID, user.Username)
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