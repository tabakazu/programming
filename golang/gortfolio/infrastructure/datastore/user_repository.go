package datastore

import (
	"github.com/jinzhu/gorm"
	"github.com/tabakazu/gortfolio/domain/entity"
)

type UserRepository struct {
	Conn *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		Conn: Config,
	}
}

func (r *UserRepository) Create(user *entity.User) error {
	if err := r.Conn.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) FindByEmail(email string) (entity.User, error) {
	var user entity.User
	if err := r.Conn.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
