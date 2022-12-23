package JWT

import (
	"log"

	"github.com/go-chi/jwtauth"
	"golang.org/x/crypto/bcrypt"
)

var TokenAuth *jwtauth.JWTAuth

// TODO: SET SECRET_KEY IN CONFIG
const SECRET_KEY = "<UMA-SENHA-MUITO-SECRETA>" // Replace <UMA-SENHA-MUITO-SECRETA> with your secret key that is private to you.

func init() {
	TokenAuth = jwtauth.New("HS256", []byte(SECRET_KEY), nil)
}

func GetHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func GenerateJWT(email string) string {
	_, tokenString, _ := TokenAuth.Encode(map[string]interface{}{"email": email})
	return tokenString
}
