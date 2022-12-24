package controller

import (
	"encoding/json"
	"net/http"
	"time"

	auth_jwt "msg-app/src/auth/jwt"
	um "msg-app/src/core/users/model"
)

func (uc *UserController) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var registerUser um.RegisterUser
	err := json.NewDecoder(r.Body).Decode(&registerUser)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	registerUser.Password = auth_jwt.GetHash([]byte(registerUser.Password))

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

	jwtToken := auth_jwt.GenerateJWT("email", loginUser.Email)

	http.SetCookie(w, &http.Cookie{
		HttpOnly: true,
		Expires:  time.Now().Add(30 * time.Second),
		SameSite: http.SameSiteLaxMode,
		// Uncomment below for HTTPS:
		// Secure: true,
		Name:  "jwt", // Must be named "jwt" or else the token cannot be searched for by jwtauth.Verifier.
		Value: jwtToken,
	})

	http.Redirect(w, r, "/users/profile", http.StatusSeeOther)
}

func (uc *UserController) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		HttpOnly: true,
		MaxAge:   -1, // Delete the cookie.
		SameSite: http.SameSiteLaxMode,
		// Uncomment below for HTTPS:
		// Secure: true,
		Name:  "jwt",
		Value: "",
	})

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
