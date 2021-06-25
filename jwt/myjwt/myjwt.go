package myjwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

// Base64Encode encodes the supplied data to Base64 string
func Base64Encode(d []byte) string {
	return base64.RawURLEncoding.EncodeToString(d)
}

// Base64Decode decodes the supplied Base64 string to original data
func Base64Decode(b64 string) []byte {
	d, err := base64.RawURLEncoding.DecodeString(b64)
	if err != nil {
		return nil
	}
	return d
}

// MakeToken creates a JWT with HS256 signature from the supplied claims
func MakeToken(claims map[string]interface{}, key string) (string, error) {
	// header
	header := make(map[string]interface{})
	header["alg"] = "HS256"
	header["typ"] = "JWT"

	// signature
	headerJSON, err := json.Marshal(header)
	if err != nil {
		return "", errors.Wrap(err, "could not create header JSON")
	}
	claimsJSON, err := json.Marshal(claims)
	if err != nil {
		return "", errors.Wrap(err, "could not create claims JSON")
	}
	b64Header := Base64Encode(headerJSON)
	b64Claims := Base64Encode(claimsJSON)
	signature := Signature(b64Header, b64Claims, key)

	// token
	token := b64Header + "." + b64Claims + "." + Base64Encode(signature)
	return token, nil
}

// Signature generates the signature of the base64 encoded header+claims with the supplied key
func Signature(b64Header string, b64Claims string, key string) []byte {
	message := b64Header + "." + b64Claims
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(message))
	return mac.Sum(nil)
}

// MapToStruct converts a map[string]interface{} to the supplied struct type
func MapToStruct(m map[string]interface{}, v interface{}) error {
	b, err := json.Marshal(m)
	if err != nil {
		return errors.Wrap(err, "failed to marshal map to json")
	}
	err = json.Unmarshal(b, v)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal struct from json")
	}
	return nil
}

// GetClaims return the claims (map) from the supplied JWT string
func GetClaims(token string, key string) (map[string]interface{}, error) {
	// Get parts
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, errors.New(fmt.Sprintf("wrong number of JWT parts: expected 3, got %v", len(parts)))
	}

	// header
	headerBytes := Base64Decode(parts[0])
	header := make(map[string]interface{})
	err := json.Unmarshal(headerBytes, &header)
	if err != nil {
		return nil, errors.Wrap(err, "problem with JWT header")
	}
	if header["alg"] != "HS256" || header["typ"] != "JWT" {
		return nil, errors.New("only support JWTs with signature of type HS256")
	}

	// claims
	claimsBytes := Base64Decode(parts[1])
	claims := make(map[string]interface{})
	err = json.Unmarshal(claimsBytes, &claims)
	if err != nil {
		return nil, errors.Wrap(err, "problem with JWT claims")
	}

	// signature check
	calculatedSignature := Base64Encode(Signature(parts[0], parts[1], key))
	signature := parts[2]
	if calculatedSignature != signature {
		return nil, errors.New(fmt.Sprintf("invalid signature: expected %v, got %v", calculatedSignature, signature))
	}

	return claims, nil
}
