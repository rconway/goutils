package jwt

import (
	"log"
	"os"
)

const defaultSigningKey = "MySecretKey"
const signingKeyEnvVar = "JWT_SIGN_KEY"

var signingKey = []byte(defaultSigningKey)

func init() {
	initSigningKey()
}

func initSigningKey() {
	key, ok := os.LookupEnv(signingKeyEnvVar)
	if !ok || len(key) == 0 {
		log.Printf("WARN: Using default signing key. Override with env '%v'\n", signingKeyEnvVar)
		signingKey = []byte(defaultSigningKey)
	} else {
		signingKey = []byte(key)
	}
	log.Printf("Using SIGN KEY = %v\n", string(signingKey))
}
