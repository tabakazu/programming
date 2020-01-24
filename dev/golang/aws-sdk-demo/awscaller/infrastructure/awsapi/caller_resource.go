package awsapi

import (
	"regexp"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/tabakazu/awscaller/model"
)

type CallerResource struct{}

func NewCallerResource() *CallerResource {
	return &CallerResource{}
}

func (a *CallerResource) GetIdentity() (*model.Caller, error) {
	svc := sts.New(session.New())
	input := &sts.GetCallerIdentityInput{}

	result, err := svc.GetCallerIdentity(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				return nil, aerr
			}
		}
		return nil, err
	}

	c := &model.Caller{
		AccountId: *result.Account,
		UserId:    *result.UserId,
		UserName:  regexp.MustCompile(`^(\S.*)\/`).ReplaceAllString(*result.Arn, ""),
	}
	return c, nil
}

func (a *CallerResource) GetAttachedIamPoliciesByUserName(userName string) (*model.Caller, error) {
	svc := iam.New(session.New())
	input := &iam.ListAttachedUserPoliciesInput{
		UserName: aws.String(userName),
	}
	result, err := svc.ListAttachedUserPolicies(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				return nil, aerr
			}
		}
		return nil, err
	}

	policies := []string{}
	for _, policy := range result.AttachedPolicies {
		policies = append(policies, *policy.PolicyName)
	}

	c := &model.Caller{
		AttachedUserPolicies: policies,
	}
	return c, nil
}
