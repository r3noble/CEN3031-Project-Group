package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	//_ "github.com/lib/pq"
)

type Credentials struct {
	Password string `json:"password", db:"password"`
	Username string `json:"username", db:"username"`
}

func Signup(w http.ResponseWriter, r *http.Request) {
	creds := &Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		//if issue occurs throw 400 bad request error
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//use bcrypt to salt & hash pword
	//8 is hashing value
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 8)

	//insert username and hashed pword into db
	if _, err = db.Query("insert into users values ($1, $2)", creds.Username, string(hashedPassword)); err != nil {
		//if error return service error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func Signin(w http.ResponseWriter, r *http.Request) {
	//create new credentials instance from parsed request
	creds := &Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//get pword found in DB for username input
	result := db.QueryRow("select password from users where username=$1", creds.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//create another credentials instance to store what we retrieve from the db
	stored := &Credentials{}
	//store obtained pword
	err = result.Scan(&stored.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			//**DO SOMETHING****
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Compare the stored hashed password, with the hashed version of the password that was received
	if err = bcrypt.CompareHashAndPassword([]byte(stored.Password), []byte(creds.Password)); err != nil {
		//***** If the two passwords don't match, return a 401 status and DO SOMETHING *****
		w.WriteHeader(http.StatusUnauthorized)
	}
}
