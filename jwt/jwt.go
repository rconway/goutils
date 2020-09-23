package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

// EncodeJSON encodes the supplied JSON object tree as a JSON string
func EncodeJSON(j map[string]interface{}) []byte {
	b, err := json.Marshal(j)
	if err != nil {
		return nil
	}
	return b
}

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
func MakeToken(claims map[string]interface{}, key string) string {
	// header
	header := make(map[string]interface{})
	header["alg"] = "HS256"
	header["typ"] = "JWT"

	// signature
	message := Base64Encode(EncodeJSON(header)) + "." + Base64Encode(EncodeJSON(claims))

	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(message))
	signature := mac.Sum(nil)

	// token
	token := message + "." + Base64Encode(signature)
	return token
}
