package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

type User struct {
	Id         int    `db:"id"`
	Username   string `db:"username"`
	Password   string `db:"password"`
	CreateTime string `db:"create_time"`
	UpdateTime string `db:"update_time"`
}

var db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
		return
	}
	db = database

}

func main() {
	var user []User
	err := db.Select(&user, "SELECT * FROM user")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	fmt.Println("insert succ:", user)
	defer db.Close()

}
