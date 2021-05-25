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
	"github.com/sandjuarezg/http-sqlite/server/database/model"
)

func main() {
	function.SqlMigration()
	var db, err = sql.Open("sqlite3", "./database/user.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var mux = http.NewServeMux()

	mux.Handle("/user/add", add(db))
	mux.Handle("/user/show", show(db))
	mux.Handle("/user/default", http.NotFoundHandler())

	fmt.Println("Listening on localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func add(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body, err = ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		err = model.AddUser(db, body)
		if err != nil {
			log.Fatal(err)
		}

		io.WriteString(w, "Data saved successfully\n")
	})
}

func show(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var text, err = model.ShowUser(db)
		if err != nil {
			log.Fatal(err)
		}

		io.WriteString(w, text)
	})
}
