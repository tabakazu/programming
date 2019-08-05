package infrastructure

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/tabakazu/interface-pattern/domain"
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
	db.Conn.AutoMigrate(&domain.Post{})
	return db
}
