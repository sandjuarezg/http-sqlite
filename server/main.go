package main

import (
	"database/sql"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sandjuarezg/http-sqlite/server/database/function"
	"github.com/sandjuarezg/http-sqlite/server/database/user"
)

func main() {
	function.SqlMigration()

	http.HandleFunc("/user/add", makeHandler(postAdd))
	http.HandleFunc("/user/show", makeHandler(getShow))
	http.HandleFunc("/user/default", http.NotFound)

	fmt.Println("Listening on localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, *sql.DB)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db, err := sql.Open("sqlite3", "./database/user.db")
		if err != nil {
			log.Fatal(err)
		}

		fn(w, r, db)
	}
}

func postAdd(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	defer db.Close()

	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		err = user.AddUser(db, body)
		if err != nil {
			log.Fatal(err)
		}

		io.WriteString(w, "Insert data successfully\n")

	}
}

func getShow(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	defer db.Close()

	if r.Method == "GET" {
		text, err := user.ShowUser(db)
		if err != nil {
			log.Fatal(err)
		}

		io.WriteString(w, text)
	}
}
