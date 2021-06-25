package jwt

import (
	"testing"
	"time"
)

// func TestMain(m *testing.M) {
// 	ExampleClaims = make(map[string]interface{})
// 	ExampleClaims["username"] = "rconway"
// 	ExampleClaims["address"] = "Stotfold"
// 	os.Exit(m.Run())
// }

func TestClaims(t *testing.T) {
	type CustomClaims struct {
		Age int `json:"age"`
		UserClaims
	}

	// Outgoing claims
	claimsTx := CustomClaims{
		Age: 123,
		UserClaims: UserClaims{
			Username: "fred",
			StandardClaims: StandardClaims{
				Issuer:    "richard",
				ExpiresAt: time.Now().Unix() + 60,
			},
		},
	}

	// Create token as string
	ss, err := ClaimsToSignedString(&claimsTx)
	t.Log(ss)

	if err != nil {
		t.Fatal(err)
	}

	// Receive token as string
	claims, err := ClaimsFromSignedString(ss, &CustomClaims{})
	claimsRx := claims.(*CustomClaims)

	if err == nil {
		t.Logf("Age: %v - Username: %v - Issuer: %v - Expires: %v\n", claimsRx.Age, claimsRx.Username, claimsRx.Issuer, claimsRx.ExpiresAt-time.Now().Unix())
	} else {
		t.Fatal("ERROR decoding token")
	}
}
