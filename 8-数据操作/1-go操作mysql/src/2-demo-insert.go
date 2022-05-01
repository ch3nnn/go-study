package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println(err)
		return
	}

	Db = database
}

func main() {
	exec, err := Db.Exec(
		"insert into test.user(username, password, create_time, update_time) values (?, ?, ?, ? )",
		"test01", "1234556", time.Now(), time.Now())
	if err != nil {
		fmt.Println(err)
	}
	insertId, err := exec.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("insert succ:", insertId)
}
