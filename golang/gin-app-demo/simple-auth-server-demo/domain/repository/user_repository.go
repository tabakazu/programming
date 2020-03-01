package repository

import "github.com/tabakazu/simple-auth-server-demo/domain"

type UserRepository interface {
	Create(email, passwdDigest, apiToken string) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	FindByToken(token string) (*domain.User, error)
}
