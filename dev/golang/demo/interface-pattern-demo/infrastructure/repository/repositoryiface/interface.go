package repositoryiface

import "github.com/tabakazu/interface-pattern-demo/domain"

type PostRepository interface {
	All() ([]domain.Post, error)
	FindById(int) (domain.Post, error)
	Store(domain.Post) (domain.Post, error)
	Update(int, domain.Post) (domain.Post, error)
	Destroy(int) error
}
