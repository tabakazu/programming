package repo

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/tabakazu/simple-auth-server-demo/domain/entry"
)

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (r *UserRepo) Create(email, passwdDigest, apiToken string) (*entry.User, error) {
	conn, err := gorm.Open("mysql", "root:@/go_simple_auth_server_dev?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	user := &entry.User{
		Email:          email,
		PasswordDigest: passwdDigest,
		APIToken:       apiToken,
	}
	if err := conn.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepo) FindByEmail(email string) (*entry.User, error) {
	conn, err := gorm.Open("mysql", "root:@/go_simple_auth_server_dev?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	user := &entry.User{}
	if err := conn.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
