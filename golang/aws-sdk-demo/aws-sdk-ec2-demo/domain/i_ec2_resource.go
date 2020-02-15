package domain

type IEC2Resource interface {
	FindById(id string) (*EC2, error)
}
