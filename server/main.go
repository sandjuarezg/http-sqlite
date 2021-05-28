//Note to self: WriteHeader must be in before writing to ResponseWriter
//When writing, it gets status 200
//if WriteHeader is called afterwards, generates a superfluous response

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

var db *sql.DB
var err error

func main() {
	function.SqlMigration()

	db, err = sql.Open("sqlite3", "./database/user.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/user/add", postAdd)
	http.HandleFunc("/user/show", getShow)

	fmt.Println("Listening on localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func postAdd(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = user.AddUser(db, body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	io.WriteString(w, "Insert data successfully\n")
}

func getShow(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	text, err := user.ShowUser(db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	io.WriteString(w, text)
}
