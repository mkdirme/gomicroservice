package auth

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

//Auth struct
type Auth struct {
	claims *Claims
	header *Header
}

//Claims for Auth
type Claims struct {
	ISS string `json:"iss"`
	SUB string `json:"sub"`
	AUD string `json:"aud"`
	IAT string `json:"iat"`
	EXP string `json:"exp"`
	JTI string `json:"jti"`
}

//Header for Auth
type Header struct {
	TYP string `json:"typ"`
	ALG string `json:"alg"`
}

//type Claims struct {
//	email string `json:"email"`
//	jwt.StandardClaims
//}

//func getToken() string{
//	mySigningKey := []byte("AllYourBase")
//	claims := Claims{"luben@aol.com",    jwt.StandardClaims{
//		ExpiresAt: 15000,
//		Issuer:    "test",
//	}}
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	ss, _ := token.SignedString(mySigningKey)
//	return ss
//}

//func decodeToken(tokenString string){
//	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//		// Don't forget to validate the alg is what you expect:
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
//		}
//
//		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
//		return []byte("my_secret_key"), nil
//	})
//
//	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
//		fmt.Println(claims["foo"], claims["nbf"])
//	} else {
//		fmt.Println(err)
//	}
//}

//NewAuth constructor for Auth
func NewAuth() *Auth {
	auth := new(Auth)
	auth.claims = new(Claims)
	auth.header = new(Header)
	return auth
}

//GenToken returns a Oauth token for a given user by email
func (auth *Auth) GenToken(email string) string {
	// claims := new(Claims)
	// header := new(Header)
	auth.claims.SUB = email
	auth.claims.AUD = "http://localhost:9002/"
	auth.claims.EXP = time.Now().AddDate(0, 0, 7).Format("2006-01-02 15:04:05")
	auth.claims.IAT = time.Now().Format("2006-01-02 15:04:05")
	auth.claims.ISS = "http://localhost:8081/"
	auth.claims.JTI = ""
	c, _ := json.Marshal(auth.claims)
	h, _ := json.Marshal(auth.header)
	str1 := base64.StdEncoding.EncodeToString(h) + "."
	str2 := base64.StdEncoding.EncodeToString(c) + "."
	accessToken := str1 + str2
	return accessToken
}

//GenerateRandom returns random int
func (auth *Auth) GenerateRandom() []byte {
	b := make([]byte, 32)
	_, _ = rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.

	return b
}

//DecodeToken recives a Oauth token and returns the Claims of the token
func (auth *Auth) DecodeToken(token string) Claims {
	var claims Claims
	if len(token) < 1 {

	}
	tokens := strings.Split(token, ".")
	if len(tokens) == 3 {
		//header, err := base64.StdEncoding.DecodeString(string(tokens[0]))
		payload, err := base64.StdEncoding.DecodeString(string(tokens[1]))
		if err != nil {
			fmt.Println("error:", err)
			//return make(byte[]),err.NewError()
		}
		json.Unmarshal(payload, &claims)
		fmt.Println(claims.EXP)
		return claims
	}

	return claims
}
