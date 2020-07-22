package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"../model"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")

	// Check if db is nil
	if db == nil {
		panic("db is nil")
	}

	// Check if err connection
	if err != nil {
		panic(err.Error())
	}

	// Check if connection successful
	if err = db.Ping(); err != nil {
		panic(err.Error())
	}

	fmt.Println("Connection Successful")

	defer db.Close()

	GetById("1")
	GetAll()

}

func GetById(id string) {
	stmt, err := db.Prepare("select email from users where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var email string
	err = stmt.QueryRow(id).Scan(&email)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(email)
}

func GetAll() {
	var users = make([]model.User, 0)

	stmt, err := db.Prepare("select * from users")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.ID, &u.Email, &u.Mobile); err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}

	for _, u := range users {
		fmt.Println(u)
	}

	//var u = user{
	//	id:     3,
	//	email:  "3@test.com",
	//	mobile: "mobile3",
	//}

	//InsertOne(u)

	GetUnknownColumns()

}

func InsertOne(u model.User) {
	stmt, err := db.Prepare("INSERT INTO users(id, email, mobile) VALUES (?,?,?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(u.ID, u.Email, u.Mobile)
	if err != nil {
		log.Fatal(err)
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Affected ", rowCnt)
}

func GetUnknownColumns() {
	stmt, err := db.Prepare("select * from users")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}

	cols, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	vals := make([]interface{}, len(cols))
	for i, _ := range cols {
		vals[i] = new(sql.RawBytes)
	}

	for rows.Next() {
		err = rows.Scan(vals...)
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, v := range vals {
		var u model.User
		fmt.Println(*(v.(*sql.RawBytes)))
		err := json.Unmarshal(*(v.(*sql.RawBytes)), &u)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(u)
	}
}
