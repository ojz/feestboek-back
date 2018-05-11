package main

import (
	"net/http"
)

func (a app) Register(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Code     string `json:"code"`
	}
	if !in(w, r, &input) {
		return
	}

	profile, err := a.repo.Register(input.Code, input.Username, input.Password)
	if err != nil {
		nok(w, 500, err)
		return
	}

	ok(w, profile)
}
