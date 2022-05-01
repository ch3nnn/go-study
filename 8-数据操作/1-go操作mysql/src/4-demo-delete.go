package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DbDelete *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println(err)
		return
	}
	DbDelete = database

}

func main() {
	res, err := DbDelete.Exec("delete from user where id=?", 56)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed, ", err)
	}
	fmt.Println("update succ:", row)
}
