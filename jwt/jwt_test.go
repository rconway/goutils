package jwt

import (
	"encoding/json"
	"strings"
	"testing"
)

const EncodedJSON = "{\"alg\":\"HS256\",\"typ\":\"JWT\"}"
const ExpectedJWT = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZGRyZXNzIjoiU3RvdGZvbGQiLCJ1c2VybmFtZSI6InJjb253YXkifQ.flmSqrq09bwTQvylFtBo4gzr-ZGmDqUtJDCLU60rIms"

func TestEncodeJSON(t *testing.T) {
	j := make(map[string]interface{})
	j["alg"] = "HS256"
	j["typ"] = "JWT"
	jbytes := EncodeJSON(j)
	jstr := string(jbytes)
	t.Log("json =", jstr)
	if string(jstr) != EncodedJSON {
		t.Error("unexpected json string:", jstr)
	}
}

func TestBase64Encode(t *testing.T) {
	b64 := Base64Encode([]byte(EncodedJSON))
	t.Log("b64 =", b64)
}

func TestMakeToken(t *testing.T) {
	// claims
	claims := make(map[string]interface{})
	claims["username"] = "rconway"
	claims["address"] = "Stotfold"

	token := MakeToken(claims, "fredbob")

	t.Log("token =", token)
	if token != ExpectedJWT {
		t.Error("wrong jwt outcome:", token)
	}
}

func TestGetClaims(t *testing.T) {
	token := ExpectedJWT

	parts := strings.Split(token, ".")

	// h := parts[0]
	c := parts[1]
	// s := parts[2]

	cb := Base64Decode(c)
	claims := make(map[string]interface{})
	err := json.Unmarshal(cb, &claims)
	if err != nil {
		t.Error("could not unmarshal claims")
	} else {
		t.Log("claims =", claims)
	}

	// convert claims to a structure
	userClaims := struct {
		Username string
		Address  string
	}{}
	b, _ := json.Marshal(claims)
	t.Log("b =", string(b))
	err = json.Unmarshal(b, &userClaims)
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log("userClaims =", userClaims)
	}
}
