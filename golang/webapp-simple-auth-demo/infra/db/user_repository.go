package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/tabakazu/simple-auth-server-demo/domain"
)

type UserRepository struct {
	Conn *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		Conn: NewDBConfig(),
	}
}

func (r *UserRepository) Create(email, passwdDigest, apiToken string) (*domain.User, error) {
	user := &domain.User{
		Email:          email,
		PasswordDigest: passwdDigest,
		APIToken:       apiToken,
	}
	if err := r.Conn.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) FindByEmail(email string) (*domain.User, error) {
	user := &domain.User{}
	if err := r.Conn.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) FindByToken(token string) (*domain.User, error) {
	user := &domain.User{}
	if err := r.Conn.Where("api_token = ?", token).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
