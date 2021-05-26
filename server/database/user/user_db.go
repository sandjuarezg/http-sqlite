package user

import (
	"bytes"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type user struct {
	id   int
	name string
	pass string
}

func AddUser(db *sql.DB, body []byte) (err error) {
	smt, err := db.Prepare("INSERT INTO users (name, password) VALUES (?, ?)")
	if err != nil {
		return
	}
	defer smt.Close()

	var element [][]byte = bytes.Split(body, []byte{'\n'})

	if len(element) == 3 {
		var user user = user{}
		user.name = string(element[0])
		user.pass = string(element[1])

		_, err = smt.Exec(user.name, user.pass)
		if err != nil {
			return
		}
	}

	return
}

func ShowUser(db *sql.DB) (text string, err error) {
	rows, err := db.Query("SELECT id, name, password FROM users")
	if err != nil {
		return
	}
	defer rows.Close()

	var user user = user{}

	text = fmt.Sprintf("|%-7s|%-15s|%-15s|\n", "id", "Name", "Password")
	text += fmt.Sprintln("_________________________________________")
	for rows.Next() {
		err = rows.Scan(&user.id, &user.name, &user.pass)
		if err != nil {
			return
		}
		text += fmt.Sprintf("|%-7d|%-15s|%-15s|\n", user.id, user.name, user.pass)
	}

	return
}
