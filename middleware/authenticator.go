package middleware

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
	"nf_stn/authentication"
)

// au instantiation
var au authentication.Auth

// TokenAuthMiddleware middleware to secure routes and authenticate requests
func TokenAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := au.TokenValid(r)
		if err != nil {
			log.Error("validation error: token invalid")
		}
		tokenAuth, err := au.ExtractTokenMetadata(r)
		if err != nil{
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
		}
		_, err = au.FetchAuth(tokenAuth)
		log.Println("token status: valid token")
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
		}

		w.Header().Add("Content-Type", "application/json")
		//w.WriteHeader(http.StatusOK)
		//info := map[string]interface{}{
		//	"authentication status": "authorized",
		//	"method":                r.Method,
		//	"content-type":          "application/json",
		//}
		//
		//encoder := json.NewEncoder(w)
		//encoder.SetIndent("", "\t")
		//_ = encoder.Encode(info)
		next(w, r)
	}
}
//
