package jwt

import (
	"log"

	jwtgo "github.com/dgrijalva/jwt-go"
)

// Alias the jwtgo types that are used on our public interfaces
type Claims = jwtgo.Claims
type StandardClaims = jwtgo.StandardClaims

// UserClaims zzz
type UserClaims struct {
	Username string `json:"username"`
	StandardClaims
}

// ToSignedString zzz
func ClaimsToSignedString(claims Claims) (string, error) {
	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, claims)
	ss, err := token.SignedString(signingKey)
	return ss, err
}

// FromSignedString zzz
func ClaimsFromSignedString(ss string, claims Claims) (Claims, error) {
	token, err := jwtgo.ParseWithClaims(ss, claims, func(token *jwtgo.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		log.Println("ERROR:", err)
		return nil, err
	} else {
		return token.Claims, nil
	}
}
