package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var secretkey []byte //use for jwt signing

func InitJWT() {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("JWT_SECRET environment variable is not set")
	}
	secretkey = []byte(secret)
}

func CreateToken(username string) (string, error) {
	//create a new jwt token for the user
	// using the secret key
	if username == "" {
		return "", fmt.Errorf("username cannot be empty")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{"username": username, //add username to the token
			"exp": time.Now().Add(time.Hour * 1).Unix(), //expires in 1 hour
			"iat": time.Now().Unix()})                   //issued at
	tokenstring, err := token.SignedString(secretkey)
	if err != nil {
		return "", err
	}
	return tokenstring, nil
}

func VerifyToken(tokenstring string) error {
	token, err := jwt.Parse(tokenstring,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return secretkey, nil
		})
	if err != nil {
		return err
	}
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}
	return nil
}
