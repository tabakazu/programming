package resource

import "github.com/tabakazu/awscaller/domain/model"

type CallerResource interface {
	GetIdentity() (*model.Caller, error)
	GetAttachedIamPoliciesByUserName(userName string) (*model.Caller, error)
}
