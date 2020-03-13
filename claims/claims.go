package claims

import (
	"os"
	"time"

	"github.com/calz10/todolister/env/constants"
	jwt "github.com/dgrijalva/jwt-go"
)

type Claims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

func GenerateNewToken(username string) string {
	superJWTKey := os.Getenv(constants.JWT_KEY)
	expirationTime := time.Now().Add(60 * 24 * time.Minute)

	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	stringToken, err := token.SignedString([]byte(superJWTKey))
	if err != nil {
		panic(err)
	}

	return stringToken
}

func VerifyToken(stringToken string) (token *jwt.Token, err bool) {

	tkn, _ := jwt.ParseWithClaims(stringToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		jwtSecretKey := []byte(os.Getenv(constants.JWT_KEY))
		return jwtSecretKey, nil
	})

	return tkn, tkn.Valid
}
