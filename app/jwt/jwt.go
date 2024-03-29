package jwt

import (
	"crud_v2/app/enviroment"
	"crud_v2/app/redis"
	"crud_v2/entity"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	uuid "github.com/nu7hatch/gouuid"
)

func getKey() []byte {
	var key []byte
	if enviroment.Get("APP_SECRET") == "" {
		key = []byte(enviroment.Get("APP_SECRET"))
	} else {
		key = []byte("secret")
	}

	return key
}

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	AuthId string `json:"auth_id"`
	UserId string `json:"user_id"`
	jwt.StandardClaims
}

// Create the Signin handler
func MakeJWT(creds *entity.User) (string, error) {
	key := getKey()
	expirationTime := time.Now().Add(10 * time.Minute)
	id := strconv.Itoa(int(creds.Id))
	v4, _ := uuid.NewV4()
	u, _ := uuid.NewV5(v4, []byte(id))
	claims := &Claims{
		UserId: id,
		AuthId: u.String(),
		StandardClaims: jwt.StandardClaims{
			Issuer:    id,
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	user, _ := json.Marshal(creds)
	s := redis.SetKey(u.String(), user, 10*time.Minute)
	if s.Err() != nil {
		fmt.Println(s.Err())
		return "", s.Err()
	}

	return tokenString, err
}

func Verify(tokenString string) (bool, error, *Claims) {
	tknStr := tokenString
	key := getKey()
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return false, err, nil
		}
		return false, err, nil
	}
	if !tkn.Valid {
		return false, err, nil
	}

	return true, nil, claims
}
