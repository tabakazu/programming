package db

import "github.com/jinzhu/gorm"

func NewDBConfig() *gorm.DB {
	conn, err := gorm.Open("mysql", "root:@/go_simple_auth_server_dev?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	return conn
}
