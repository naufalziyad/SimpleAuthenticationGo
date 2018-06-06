package main

// Author: Naufal Ziyad Luthfiansyah //

import ("database/sql"
		"fmt"
		/*"html/template"*/
		"log"
		"net/http"

		"golang.org/x/crypto/bcrypt"

		_ "github.com/go-sql-driver/mysql"
		/*"github.com/kataras/go-sessions"*/)

var db *sql.DB
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

	defer db.Close()

	fmt.Println("Server running on port :8000")
	http.ListenAndServe(":8000", nil)
}

func QueryUser(username string) user {
	var users = user{}
	err = db.QueryRow(`
		SELECT id,
		username,
		first_name,
		last_name,
		password
		FROM users WHERE username=?
		`, username).
		Scan(
			&users.ID,
			&users.Username,
			&users.Firstname,
			&users.Lastname,
			&users.Password,
			)
	return users
}

func register(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST" {
		http.ServeFile(w, r, "register.html")
		return
	}

	username :=r.FormValue("email")
	first_name :=r.FormValue("first_name")
	last_name :=r.FormValue("last_name")
	password :=r.FormValue("password")

	users := QueryUser(username)

	if (user{}) ==users {
		hashedPassword, err  := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

		if len(hashedPassword) !=0 && checkErr (w, r, err) {
			stmt, err := db.Prepare("INSERT INTO users SET username=?, password=?, first_name=?, last_name=? ")
			if err == nil {
				_, err := stmt.Exec(&username, &hashedPassword, &first_name, &last_name)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return

			}
		} else {
			http.Redirect(w, r, "/register",302)
		}
	}
}

func checkErr(w http.ResponseWriter, r *http.Request, err error) bool {
 if err != nil {

fmt.Println(r.Host + r.URL.Path)

http.Redirect(w, r, r.Host+r.URL.Path, 301)
 return false
 }

return true
}