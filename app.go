package main

import (
	"log"
	"net/http"
	"strings"
)

type app struct {
	config  config
	repo    *repo
	handler *http.ServeMux
}

func build(c config) *app {
	a := &app{}

	a.config = c
	a.repo = buildRepo(c)

	r := http.NewServeMux()
	r.HandleFunc("/api/login", a.Login)
	r.HandleFunc("/api/logout", a.Logout)
	a.handler = r

	return a
}

func (a app) run() {
	var url string
	if strings.HasPrefix(a.config.address, ":") {
		url = "http://localhost" + a.config.address
	} else {
		url = "http://" + a.config.address
	}

	log.Println("Launching server on " + url)
	log.Fatal(http.ListenAndServe(a.config.address, a.handler))
}
