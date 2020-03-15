package datastore

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Config struct {
	Conn *gorm.DB
}

func NewConfig() *Config {
	conn, err := gorm.Open("mysql", "root:@/domain_modeling_demo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}

	config := new(Config)
	config.Conn = conn
	return config
}
