package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/tabakazu/interface-pattern/interfaces/http"
)

type DB struct {
	Conn *gorm.DB
}

func NewDb() *DB {
	conn, err := gorm.Open("mysql", "root:@/simple_rest_api_dev?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	db := new(DB)
	db.Conn = conn
	db.Conn.AutoMigrate(&Post{})
	return db
}

type Post struct {
	gorm.Model
	Title string `gorm:"not null" form:"title" json:"title"`
	Body  string `gorm:"not null" form:"body" json:"body"`
}

// type PostRepository interface {
// 	FindByID(postID int) (*Post, error)
// }

// type PostRepositoryDao struct {
// 	*DB
// }

// func (PostRepositoryDao) FindByID(postID int) (*Post, error) {
// 	post := Post{}
// 	if err := conn.Find(&post, "id = ?", postID).Error; err != nil {
// 		return nil, err
// 	}
// 	return &post, nil
// }

func main() {
	db := NewDb()
	conn := db.Conn
	fmt.Println(conn)

	fmt.Println("Hello Interface Pattern")

	r := http.NewRouter()
	r.Run(":8080")
}
