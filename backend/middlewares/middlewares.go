package middlewares

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JwtKey *ecdsa.PrivateKey
var JwtPublicKey *ecdsa.PublicKey

func generateJwtKey() {
	jwtKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	JwtKey = jwtKey
	JwtPublicKey = &jwtKey.PublicKey
}

type Claims struct {
	PersonId int64 `json:"personId"`
	jwt.StandardClaims
}

func CreateToken(w http.ResponseWriter, personId int64) (err error) {
	generateJwtKey()
	claims := &Claims{
		PersonId: personId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		panic(err)
		return err
	}

	http.SetCookie(w,
		&http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: time.Now().Add(time.Hour * 24),
		},
	)
	return nil

}

func VerifyToken(r *http.Request) (int64, error) {

	cookie, err := r.Cookie("token")
	if err != nil {
		return -1, err
	}
	tokenString := cookie.Value
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtPublicKey, nil
	})

	if err != nil {
		fmt.Println("invalid jwtKey")
		return -1, err
	}

	if !token.Valid {
		return -1, fmt.Errorf("invalid token")
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		// log.Printf("%v", claims.PersonId)
		// fmt.Printf("%v", claims.PersonId)
		return claims.PersonId, nil
	}

	return -1, nil
}
