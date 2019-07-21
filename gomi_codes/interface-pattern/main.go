package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Db struct {
	Conn *gorm.DB
}

func NewDb() *Db {
	conn, err := gorm.Open("mysql", "root:@/simple_rest_api_dev?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	db := new(Db)
	db.Conn = conn
	db.Conn.AutoMigrate(&Post{})
	return db
}

type Post struct {
	gorm.Model
	Title string `gorm:"not null" form:"title" json:"title"`
	Body  string `gorm:"not null" form:"body" json:"body"`
}

func main() {
	db := NewDb()
	conn := db.Conn
	fmt.Println(conn)

	fmt.Println("Hello Interface Pattern")
}
