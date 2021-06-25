package myjwt

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

const ExpectedBase64 = "eyJhZGRyZXNzIjoiU3RvdGZvbGQiLCJ1c2VybmFtZSI6InJjb253YXkifQ"
const ExpectedJWT = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZGRyZXNzIjoiU3RvdGZvbGQiLCJ1c2VybmFtZSI6InJjb253YXkifQ.flmSqrq09bwTQvylFtBo4gzr-ZGmDqUtJDCLU60rIms"
const Key = "fredbob"

var ExampleClaims map[string]interface{}

func TestMain(m *testing.M) {
	ExampleClaims = make(map[string]interface{})
	ExampleClaims["username"] = "rconway"
	ExampleClaims["address"] = "Stotfold"
	os.Exit(m.Run())
}

func TestBase64Encode(t *testing.T) {
	sampleJSON, err := json.Marshal(ExampleClaims)
	if err != nil {
		t.Error("error creating ExampleClaims as JSON:", err)
	}
	b64 := Base64Encode([]byte(sampleJSON))

	if b64 != ExpectedBase64 {
		t.Error("wrong base64 outcome:", b64)
	} else {
		t.Log("ExampleClaims (base64) =", b64)
	}
}

func TestMakeToken(t *testing.T) {
	// claims
	claims := ExampleClaims

	// create token
	token, err := MakeToken(claims, Key)
	if err != nil {
		t.Error("error making token:", err)
	}

	// check
	if token != ExpectedJWT {
		t.Error("wrong jwt outcome:", token)
	} else {
		t.Log("token =", token)
	}
}

func TestGetClaims(t *testing.T) {
	claims, err := GetClaims(ExpectedJWT, Key)
	if err != nil {
		t.Error("error getting claims from token:", err)
	}

	for k, v := range ExampleClaims {
		if v != claims[k] {
			t.Error(fmt.Sprintf("claim check FAILURE [%v]: expected %v, got %v", k, v, claims[k]))
		} else {
			t.Log(fmt.Sprintf("claim check SUCCESS [%v]: %v", k, v))
		}
	}
}
