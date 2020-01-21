package infrastructure

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DbConfig struct {
	Conn *gorm.DB
}

func NewDbConfig() DbConfig {
	conn, err := gorm.Open("mysql", "root:@/golang_api_srv_dev?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	return DbConfig{Conn: conn}
}
