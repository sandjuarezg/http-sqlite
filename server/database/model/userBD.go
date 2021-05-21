package model

import (
	"bytes"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func AddUser(db *sql.DB, body []byte) (err error) {
	smt, err := db.Prepare("INSERT INTO users (name, password) VALUES (?, ?)")
	if err != nil {
		return
	}
	defer smt.Close()

	var element = bytes.Split(body, []byte("\n"))

	_, err = smt.Exec(element[0], element[1])
	if err != nil {
		return
	}

	return
}

func ShowUser(db *sql.DB) (text string, err error) {
	rows, err := db.Query("SELECT id, name, password FROM users")
	if err != nil {
		return
	}
	defer rows.Close()

	var element [3]string

	text = fmt.Sprintf("|%-7s|%-15s|%-15s|\n", "id", "Name", "Password")
	text += fmt.Sprintln("_________________________________________")
	for rows.Next() {
		err = rows.Scan(&element[0], &element[1], &element[2])
		if err != nil {
			return
		}
		text += fmt.Sprintf("|%-7s|%-15s|%-15s|\n", element[0], element[1], element[2])
	}

	return
}
