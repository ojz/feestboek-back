package main

import (
	"net/http"
)

func (a app) Login(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if !in(w, r, &input) {
		return
	}

	// byteHash := []byte(hashedPwd)

	// err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	// if err != nil {
	//     log.Println(err)
	//     return false
	// }

	// return true

	profile, err := a.repo.Login(input.Username, input.Password)
	if err != nil {
		nok(w, 500, err)
		return
	}

	ok(w, profile)
}
