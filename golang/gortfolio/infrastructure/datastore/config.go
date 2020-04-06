package datastore

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Config *gorm.DB

func init() {
	// conn, err := gorm.Open("mysql", "root:root@/gortfolio_dev?charset=utf8&parseTime=True&loc=Local")
	conn, err := gorm.Open("mysql", "root:root@tcp(mysql:3306)/gortfolio_dev?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err.Error())
	}
	Config = conn
}
