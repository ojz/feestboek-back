package main

import (
	"database/sql"
	"errors"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (a app) Login(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if !in(w, r, &input) {
		return
	}

	profile, err := a.repo.Login(input.Username, input.Password)
	if err != nil {
		nok(w, 500, err)
		return
	}

	ok(w, profile)
}

func (s repo) Login(username, password string) (*Profile, error) {
	var user struct {
		ID        int64
		Hash, Bio string
	}
	err := s.getUser.Get(&user, username)
	if err == sql.ErrNoRows {
		return nil, errors.New("User not found.")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(password))
	if err != nil {
		return nil, errors.New("Invalid password.")
	}

	profile := Profile{
		ID:       user.ID,
		Username: username,
		Bio:      user.Bio,
	}
	return &profile, nil
}
