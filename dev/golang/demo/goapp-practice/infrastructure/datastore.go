package infrastructure

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/tabakazu/goapp-practice/domain/model"
)

type Datastore struct {
	Session *gorm.DB
}

func NewDatastore() *Datastore {
	sess, err := gorm.Open("mysql", "root:@/goapp_practice_dev?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	db := new(Datastore)
	db.Session = sess
	db.Session.AutoMigrate(&model.Post{})
	return db
}
