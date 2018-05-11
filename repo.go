package main

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

type repo struct {
	db *sqlx.DB

	useCode, getInvitation, addUser,
	getUser *sqlx.Stmt
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
		getInvitation: prepare(`SELECT id FROM invitations WHERE code = ?`),
		useCode:       prepare(`UPDATE invitations SET used = 1 WHERE used = 0 AND id = ?`),
		addUser:       prepare(`INSERT INTO users (username, hash, invitation) VALUES (?, ?, ?)`),
		getUser:       prepare(`SELECT id, hash, bio FROM users WHERE username = ?`),
	}
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
