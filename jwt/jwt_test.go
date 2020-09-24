package jwt

import (
	"fmt"
	"os"
	"testing"
)

const EncodedJSON = "{\"alg\":\"HS256\",\"typ\":\"JWT\"}"
const ExpectedJWT = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZGRyZXNzIjoiU3RvdGZvbGQiLCJ1c2VybmFtZSI6InJjb253YXkifQ.flmSqrq09bwTQvylFtBo4gzr-ZGmDqUtJDCLU60rIms"
const Key = "fredbob"

var ExampleClaims map[string]interface{}

func TestMain(m *testing.M) {
	ExampleClaims = make(map[string]interface{})
	ExampleClaims["username"] = "rconway"
	ExampleClaims["address"] = "Stotfold"
	os.Exit(m.Run())
}

func TestEncodeJSON(t *testing.T) {
	j := make(map[string]interface{})
	j["alg"] = "HS256"
	j["typ"] = "JWT"
	jbytes, err := EncodeJSON(j)
	if err != nil {
		t.Error(err)
	}
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
	claims := ExampleClaims

	token, err := MakeToken(claims, Key)
	if err != nil {
		t.Error(err)
	}

	t.Log("token =", token)
	if token != ExpectedJWT {
		t.Error("wrong jwt outcome:", token)
	}
}

func TestGetClaims(t *testing.T) {
	claims, err := GetClaims(ExpectedJWT, Key)
	if err != nil {
		t.Error("error getting claims from token")
	}

	if claims["username"] != "rconway" {
		t.Error(fmt.Sprintf("wrong username: expected rconway, got %v", claims["username"]))
	}

	t.Log("claims =", claims)
}
