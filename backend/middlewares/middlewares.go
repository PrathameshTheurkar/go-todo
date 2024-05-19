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
	Username string `json:"username"`
	jwt.StandardClaims
}

func CreateToken(w http.ResponseWriter, username string) (err error) {
	generateJwtKey()
	claims := &Claims{
		Username: username,
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

func VerifyToken(r *http.Request) error {

	cookie, err := r.Cookie("token")
	if err != nil {
		return err
	}
	tokenString := cookie.Value
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtPublicKey, nil
	})
	if err != nil {
		fmt.Println("invalid jwtKey")
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}
	fmt.Println(token)

	return nil
}
