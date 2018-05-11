package main

type Profile struct {
	ID       int64  `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Bio      string `db:"bio" json:"bio"`
}
