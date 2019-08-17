package usecase

type Post interface {
	FindByID(int) error
}
