package model

type Caller struct {
	AccountId            string
	UserId               string
	UserName             string
	AttachedUserPolicies []string
}

type CallerResource interface {
	GetIdentity() (*Caller, error)
	GetAttachedIamPoliciesByUserName(userName string) (*Caller, error)
}
