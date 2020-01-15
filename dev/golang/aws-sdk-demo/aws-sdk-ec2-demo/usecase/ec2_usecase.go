package usecase

import "github.com/tabakazu/aws-sdk-ec2-demo/domain"

type EC2Usecase struct {
	Resource domain.IEC2Resource
}

func (u *EC2Usecase) GetById(id string) (*domain.EC2, error) {
	ec2, err := u.Resource.FindById(id)
	if err != nil {
		return nil, err
	}
	return ec2, nil
}
