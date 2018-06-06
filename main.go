package main

import ("database/sql"
		"fmt"
		"html/template"
		"log"
		"net/http"

		"golang.org/x/crypto/bcrypt"

		_ "github.com/go-sql-driver/mysql"
		"github.com/kataras/go-sessions")

var db *sql.db
var err error

type user struct {
	ID 			int
	Username	string
	Firstname	string
	Lastname	string
	Password	string
}

func connect_db(){
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1)/go_db")

	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
}

func routes(){
	http.HandleFunc("/register", register)
}

func main(){
	connect_db()
	routes()

	defer db.Closer()

	fmt.Println("Server running on port :8000")
	http.ListenAndServe(":8000", nil)
}

func QueryUser(username string) user {
	var users = user{}
	err = db.QueryRow('
		SELECT id,
		username,
		first_name,
		last_name,
		password
		FROM users WHERE username=?
		', username).
		Scan(
			&users.ID,
			&users.Username,
			&users.Firstname,
			&users.Lastname,
			&users.Password,
			)
	return users
}

func register(w.http.ResponseWriter, r *http.Request){
	if r.Method != "POST" {
		http.ServeFile(w, r, "register.html")
		return
	}



}