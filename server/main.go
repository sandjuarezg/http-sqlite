package main

import (
	"database/sql"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	function "github.com/sandjuarezg/http-sqlite/server/database/functionality"
	"github.com/sandjuarezg/http-sqlite/server/database/model"
)

func user() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "1. Add user\n")
		io.WriteString(w, "2. Show users\n")

		var body, err = ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		if len(body) != 0 {
			body = body[:len(body)-1]
			fmt.Printf("client says: %s\n", body)

			switch string(body) {
			case "1":
				fmt.Fprintln(w, "/add")
			case "2":
				fmt.Fprintln(w, "/show")
			default:
				fmt.Fprintln(w, "/default")
			}
		}
	})
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

func main() {
	function.SqlMigration()
	var db, err = sql.Open("sqlite3", "./database/user.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var mux = http.NewServeMux()

	mux.Handle("/user", user())
	mux.Handle("/user/add", add(db))
	mux.Handle("/user/show", show(db))
	mux.Handle("/user/default", http.NotFoundHandler())

	fmt.Println("Listening on localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", mux))
}
