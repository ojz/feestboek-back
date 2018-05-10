package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type repo struct {
	db            *sqlx.DB
	checkPassword *sqlx.Stmt
}

func buildRepo(c config) *repo {
	db := sqlx.MustConnect("sqlite3", c.database)
	prepare := func(sql string) *sqlx.Stmt {
		stmt, err := db.Preparex(sql)
		if err != nil {
			panic(err)
		}
		return stmt
	}

	return &repo{
		checkPassword: prepare(`SELECT id, username, bio FROM users WHERE username = ? and password = ?`),
	}
}

func (s repo) Login(username, password string) (*Profile, error) {
	profile := Profile{}
	err := s.checkPassword.Get(&profile, username, password)
	if err != nil {
		return nil, err
	}

	return &profile, nil
}
