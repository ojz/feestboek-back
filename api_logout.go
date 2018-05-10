package main

import "net/http"

func (a app) Logout(w http.ResponseWriter, r *http.Request) {
	ok(w, nil)
}
