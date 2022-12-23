package controller

import (
	"encoding/json"
	"log"
	"net/http"

	um "msg-app/internal/core/users/model"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// TODO NEW PASS
var SECRET_KEY = []byte("gosecretkey")

func (uc *UserController) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var registerUser um.RegisterUser
	err := json.NewDecoder(r.Body).Decode(&registerUser)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	registerUser.Password = getHash([]byte(registerUser.Password))

	if err := uc.service.Register(registerUser); err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(`{"message":"User created."}`))
}

func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var loginUser um.LoginUser
	err := json.NewDecoder(r.Body).Decode(&loginUser)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	ok := uc.service.Login(loginUser)
	if !ok {
		w.Write([]byte(`{"message":"Wrong email or password."}`))
		return
	}

	jwtToken, err := generateJWT()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}
	w.Write([]byte(`{"token":"` + jwtToken + `"}`))
}

func getHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func generateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString(SECRET_KEY)
	if err != nil {
		log.Println("Error in JWT token generation")
		return "", err
	}
	return tokenString, nil
}
