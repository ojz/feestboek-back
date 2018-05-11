package main

import (
	"fmt"
	"net/http"
)

func (a app) Deregister(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if !in(w, r, &input) {
		return
	}

	err := a.repo.Deregister(input.Username, input.Password)
	if err != nil {
		nok(w, 500, err)
		return
	}

	ok(w, nil)
}

func (s repo) Deregister(username, password string) error {
	_, err := s.Login(username, password)
	if err != nil {
		return err
	}

	res, err := s.dropUser.Exec(username)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count != 1 {
		// panic?
		return fmt.Errorf("%v users deleted.", count)
	}

	return nil
}
