package infrastructure

import (
	"github.com/jinzhu/gorm"
)

type Db struct {
	Conn *gorm.DB
}
