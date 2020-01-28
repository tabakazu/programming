package usecase

import (
	"github.com/tabakazu/awscaller/domain/resource"
	"github.com/tabakazu/awscaller/infrastructure/awsapi"
)

type CallerUsecase struct {
	Resource resource.CallerResource
}

func NewCallerUsecase() *CallerUsecase {
	return &CallerUsecase{
		Resource: awsapi.NewCallerResource(),
	}
}

func (u *CallerUsecase) GetAccountId() (string, error) {
	r := awsapi.NewCallerResource()
	clr, err := r.GetIdentity()
	if err != nil {
		return "", err
	}
	return clr.AccountId, nil
}

func (u *CallerUsecase) GetUserId() (string, error) {
	r := awsapi.NewCallerResource()
	clr, err := r.GetIdentity()
	if err != nil {
		return "", err
	}
	return clr.UserId, nil
}

func (u *CallerUsecase) GetUserName() (string, error) {
	r := awsapi.NewCallerResource()
	clr, err := r.GetIdentity()
	if err != nil {
		return "", err
	}
	return clr.UserName, nil
}

func (u *CallerUsecase) GetAttachedUserPolicies(userName string) ([]string, error) {
	r := awsapi.NewCallerResource()
	clr, err := r.GetAttachedIamPoliciesByUserName(userName)
	if err != nil {
		return nil, err
	}
	return clr.AttachedUserPolicies, nil
}
