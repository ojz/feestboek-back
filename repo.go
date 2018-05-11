package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type repo struct {
	db *sqlx.DB

	useCode, getInvitation, addUser,
	getUser, dropUser *sqlx.Stmt
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
		getInvitation: prepare(`SELECT id FROM invitations WHERE code = ? and used = 0`),
		useCode:       prepare(`UPDATE invitations SET used = 1 WHERE used = 0 AND id = ?`),
		addUser:       prepare(`INSERT INTO users (username, hash, invitation) VALUES (?, ?, ?)`),
		getUser:       prepare(`SELECT id, hash, bio FROM users WHERE username = ?`),
		dropUser:      prepare(`DELETE FROM users WHERE username = ?`),
	}
}
