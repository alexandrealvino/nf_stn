package authentication

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"github.com/twinj/uuid"
	"nf_stn/config"
	"os"
	"strconv"
	"strings"
	"time"

	"net/http"
	"nf_stn/entities"
)

//go:generate  go run github.com/golang/mock/mockgen  -package mock -destination=./mock/token_mock.go -source=$GOFILE

// Token interface
type Token interface {
	Init()
	ExtractToken(r *http.Request) string
	VerifyToken(r *http.Request) (*jwt.Token, error)
	TokenValid(r *http.Request) error
	ExtractTokenMetadata(r *http.Request) (*entities.AccessDetails, error)
	Logout(w http.ResponseWriter, r *http.Request)
	FetchAuth(authD *entities.AccessDetails) (uint64, error)
	CreateToken(userid uint64, username string) (*entities.TokenDetails, error)
	CreateAuth(userID uint64, td *entities.TokenDetails) error
	DeleteAuth(givenUUID string) (int64, error)
	RefreshToken(w http.ResponseWriter, r *http.Request)
}

// Auth struct
type Auth struct {
	Token
	RedisCg config.RedisConfig
}

// redis instantiation
var redis config.App

// Init initializes redis connection
func (au *Auth) Init() {
	redis.ConnectRedis(au.RedisCg.DSN())
}

// ExtractToken extracts the token from the request header
func (au *Auth) ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

// VerifyToken verifies the token
func (au *Auth) VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := au.ExtractToken(r)
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

// TokenValid checks the validity of the token, returns error if it has already expired
func (au *Auth) TokenValid(r *http.Request) error {
	token, err := au.VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

// ExtractTokenMetadata the token metadata so we can lookup in our redis
func (au *Auth) ExtractTokenMetadata(r *http.Request) (*entities.AccessDetails, error) {
	token, err := au.VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUUID, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userID, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &entities.AccessDetails{
			AccessUUID: accessUUID,
			UserID:     userID,
		}, nil
	}
	return nil, err
}

// Logout requests a logout
func (au *Auth) Logout(w http.ResponseWriter, r *http.Request) {
	tokenAuth, err := au.ExtractTokenMetadata(r)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "\t")
		_ = encoder.Encode("unauthorized")
		return
	}
	deleted, delErr := au.DeleteAuth(tokenAuth.AccessUUID)
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

// FetchAuth fetches access details from the token in redis
func (au *Auth) FetchAuth(authD *entities.AccessDetails) (uint64, error) {
	ID, err := redis.Clt.Get(authD.AccessUUID).Result()
	//ID, err := config.Client.Get(authD.AccessUUID).Result()
	if err != nil {
		return 0, err
	}
	userID, _ := strconv.ParseUint(ID, 10, 64)
	return userID, nil
}

// CreateToken creates token
func (au *Auth) CreateToken(userid uint64, username string) (*entities.TokenDetails, error) {
	var err error
	td := &entities.TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUUID = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUUID = uuid.NewV4().String()
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUUID
	atClaims["user_id"] = userid
	atClaims["username"] = username
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}
	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUUID
	rtClaims["user_id"] = userid
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}
	return td, nil
}

// CreateAuth creates authentication access
func (au *Auth) CreateAuth(userID uint64, td *entities.TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()
	errAccess := redis.Clt.Set(td.AccessUUID, strconv.Itoa(int(userID)), at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := redis.Clt.Set(td.RefreshUUID, strconv.Itoa(int(userID)), rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}

// DeleteAuth deletes authentication access
func (au *Auth) DeleteAuth(givenUUID string) (int64, error) {
	deleted, err := redis.Clt.Del(givenUUID).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}

// RefreshToken refreshes the given token
func (au *Auth) RefreshToken(w http.ResponseWriter, r *http.Request) {
	mapToken := map[string]string{}
	bearToken := r.Header.Get("Authorization")
	username := r.FormValue("username")
	mapToken["refresh_token"] = bearToken
	refreshToken := mapToken["refresh_token"]
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		refreshToken = strArr[1]
	}

	//verify the token
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})
	//if there is an error, the token must have expired
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "\t")
		_ = encoder.Encode("refresh token expired")
		return
	}
	//is token valid?
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		log.Error(err)
	}
	//Since token is valid, get the uuid:
	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
	if ok && token.Valid {
		refreshUuid, ok := claims["refresh_uuid"].(string) //convert the interface to string
		if !ok {
			w.WriteHeader(http.StatusUnprocessableEntity)
			log.Error(err)
			return
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			log.Error(err)
			return
		}
		//Delete the previous Refresh Token
		deleted, delErr := au.DeleteAuth(refreshUuid)
		if delErr != nil || deleted == 0 { //if any goes wrong
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			info := map[string]string{
				"authentication status": "unauthorized",
				"method":                r.Method,
				"content-type":          "application/json",
			}
			encoder := json.NewEncoder(w)
			encoder.SetIndent("", "\t")
			_ = encoder.Encode(info)
			log.Error(err)
			return
		}
		//Create new pairs of refresh and access tokens
		ts, createErr := au.CreateToken(userId,username)
		if  createErr != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			log.Error(createErr.Error())
			return
		}
		//save the tokens metadata to redis
		saveErr := au.CreateAuth(userId, ts)
		if saveErr != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			log.Error(saveErr.Error())
			return
		}
		tokens := map[string]string{
			"access_token":  ts.AccessToken,
			"refresh_token": ts.RefreshToken,
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "\t")
		_ = encoder.Encode(tokens)
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "\t")
		_ = encoder.Encode("refresh expired")
	}
}
