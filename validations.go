package main

import "regexp"

var validName = regexp.MustCompile(`^[a-zA-Z0-9 -_@!.]*$`)
var validPassword = regexp.MustCompile(`^[a-zA-Z0-9 -_@!.]*$`)
var validID = regexp.MustCompile(`^[0-9]{1,20}$`)
