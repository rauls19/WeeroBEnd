package utils

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//GenerateToken is a function to generate a token given an user and a SECRET key
func GenerateToken(clientid string, pass string) string {
	var jwtKey = []byte("my_secret_key")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["client"] = clientid
	claims["pass"] = pass
	claims["scope"] = "first scope think"
	claims["admin"] = false
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	tokenString, _ := token.SignedString(jwtKey)
	return tokenString
}

//GenerateRefreshToken is a function to generate a refresh token given an user and a SECRET key
func GenerateRefreshToken(sub string) string {
	var jwtKey = []byte("my_secret_key")
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = sub
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token, _ := refreshToken.SignedString(jwtKey)
	return token
}

//TokenIsValid is the function to check if the token is valid
func TokenIsValid(r *http.Request) (bool, []string, error) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]
	token, err := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["err"])
		}
		return []byte("my_secret_key"), nil
	})
	if ok := token.Valid; !ok {
		return false, nil, err
	}
	claims := token.Claims.(jwt.MapClaims)
	var claim []string
	claim[1] = claims["client"].(string)
	claim[2] = claims["pass"].(string)
	claim[3] = claims["scope"].(string)
	claim[4] = claims["sub"].(string)
	return true, claim, err
}

func GenerateClientID(userid int) string {
	hash := sha256.New()
	hash.Write([]byte(strconv.Itoa(userid)))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func GenerateCLientSecret(userid int, password string) string {
	hash := sha256.New()
	hash.Write([]byte(strconv.Itoa(userid) + password))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
