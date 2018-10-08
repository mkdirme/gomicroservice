package auth

import (
	"fmt"
	"testing"
)

func TestDecodeJWT(t *testing.T) {
	jwts := GenToken("temp@aol.com")
	claim := DecodeToken(jwts)
	fmt.Println("-L- email result for: " + claim.SUB)
}
