package main

import (
	"fmt"
	"log"
	"strings"

	"../model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

func main() {
	db, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	GetAll()

	GetById("1")

	InsertOrUpdate(model.User{
		ID:     5,
		Email:  "5@test.com",
		Mobile: "mobile5",
	})

	//GetAll()
	GetAllRaw()

}

func GetAll() {
	var users = make([]model.User, 0)
	if err := db.Table("users").Find(&users).Error; err != nil {
		log.Fatal(err)
	}
	for _, u := range users {
		fmt.Println(u)
	}
}

func GetById(id string) {
	var user model.User
	if err := db.Table("users").Where("id = ?", id).First(&user).Error; err != nil {
		log.Fatal(err)
	}

	fmt.Println(user)
}

func InsertOrUpdate(user model.User) {
	if err := db.Table("users").Save(&user).Error; err != nil {
		log.Fatal(err)
	}
}

func GetAllRaw() {
	var users = make([]model.User, 0)
	var query strings.Builder
	query.WriteString("SELECT * FROM test.users")

	err := db.Raw(query.String()).Find(&users).Error
	if gorm.IsRecordNotFoundError(err) {
		fmt.Println("not found")
		return
	}

	if err != nil {
		log.Fatal(err)
	}

	for _, u := range users {
		fmt.Println(u)
	}

}
