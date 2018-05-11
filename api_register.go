package main

import (
	"database/sql"
	"errors"
	"net/http"

	"golang.org/x/crypto/bcrypt"
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

func (r repo) Register(code, username, password string) (*Profile, error) {
	// get the invitation
	var invitation struct{ ID string }
	err := r.getInvitation.Get(&invitation, code)
	if err == sql.ErrNoRows {
		return nil, errors.New("Invalid code")
	}
	if err != nil {
		return nil, err
	}

	// check if username taken
	var user struct {
		ID        sql.NullInt64
		Bio, Hash sql.NullString
	}
	err = r.getUser.Get(&user, username)
	if err == nil {
		return nil, errors.New("Username already taken.")
	}
	if err != sql.ErrNoRows {
		return nil, err
	}

	// check & invalidate the code
	res, err := r.useCode.Exec(invitation.ID)
	if err != nil {
		return nil, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if count != 1 {
		return nil, errors.New("Invalid code.")
	}

	// hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}

	// add the user
	res, err = r.addUser.Exec(username, string(hash), invitation.ID)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	// return the profile
	profile := Profile{
		ID:       id,
		Username: username,
		Bio:      string(hash),
	}
	return &profile, nil
}
