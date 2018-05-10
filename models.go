package main

type Profile struct {
	ID       string `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Bio      string `db:"bio" json:"bio"`
}
