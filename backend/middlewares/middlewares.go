package middlewares

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = "secret-key"
var jwtKey = []byte(secretKey)

// func generateSecretKey() *ecdsa.PrivateKey {
// 	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

// 	if err != nil {
// 		panic(err)
// 	}
// 	return key
// }

// var SecretKey *ecdsa.PrivateKey

// func putSecretKey() {
// 	SecretKey = generateSecretKey()
// }

func CreateToken(username string) (string, error) {
	// secretKey := generateSecretKey()
	// fmt.Println(secretKey)
	// SecretKey = secretKey

	token := jwt.NewWithClaims(jwt.SigningMethodES256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	// secretKey := generateSecretKey()
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// fmt.Println(SecretKey)
		return jwtKey, nil
	})
	if err != nil {
		// panic(err)
		return err
	}

	// if !token.Valid {
	// 	return fmt.Errorf("invalid token")
	// }
	fmt.Println(token)

	return nil
}
