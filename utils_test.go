package main

import (
	"database/sql"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

//go:generate include -o internal/assets/testdata sql/testdata.sql
// => can call testdata.Get()

func wantPanic(t *testing.T) {
	err := recover()
	if err == nil {
		t.Fail()
	}
}

func mockDatabase(name string) string {
	return "file:" + name + "?mode=memory&cache=shared"
}

func mockDB(name string) *sql.DB {
	db, err := sql.Open("sqlite3", mockDatabase(name))
	if err != nil {
		panic(err)
	}
	return db
}

func mockStatements(name string) *statements {
	loadSchema(mockDB(name))
	loadTestdata(mockDB(name))
	return _buildSQLStatements(mockDatabase(name))
}

func loadSchema(db *sql.DB) {
	schema, err := ioutil.ReadFile("./schema.sql")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(string(schema))
	if err != nil {
		panic(err)
	}
}

func loadTestdata(db *sql.DB) {
	testdata, err := ioutil.ReadFile("./testdata.sql")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(string(testdata))
	if err != nil {
		panic(err)
	}
}

type mockTpl struct {
	name string
}

func (m *mockTpl) ExecuteTemplate(wr io.Writer, name string, data interface{}) error {
	if name != m.name {
		return errors.New("Oh")
	}
	return nil
}

func recordGET(url string, handler http.Handler) *httptest.ResponseRecorder {
	req := httptest.NewRequest("GET", url, nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return rr
}

func recordPOST(url string, data url.Values, handler http.Handler) *httptest.ResponseRecorder {
	req := httptest.NewRequest("POST", url, strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return rr
}
