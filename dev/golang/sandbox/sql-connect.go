package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@/go_sandbox?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	_, err = db.Exec(
		`CREATE TABLE posts (id INTEGER, title VARCHAR(255))`,
	)
	if err != nil {
		fmt.Println("Faild!")
	}

	fmt.Println("Successful!")

	defer db.Close()
}
