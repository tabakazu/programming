package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/iam/iamiface"
)

func main() {
	sess := iam.New(session.New())

	iamPath := "/"
	iamPermissionsBoundary := "testtesttesttesttest"

	iamTagKey := "Name"
	iamTagValue := "test-go-sdk"
	iamUserName := "test-go-sdk"
	iamTag := iam.Tag{
		Key:   &iamTagKey,
		Value: &iamTagValue,
	}
	iamTags := []*iam.Tag{&iamTag}

	user := &User{
		Client:              sess,
		Path:                iamPath,
		PermissionsBoundary: iamPermissionsBoundary,
		UserName:            iamUserName,
		Tags:                iamTags,
	}

	resp, err := user.Create()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp)
	return
}

type User struct {
	Client              iamiface.IAMAPI
	Path                string
	PermissionsBoundary string
	UserName            string
	Tags                []*iam.Tag
}

func (u *User) Create() (*iam.CreateUserOutput, error) {
	param := iam.CreateUserInput{
		Path:                &u.Path,
		PermissionsBoundary: &u.PermissionsBoundary,
		UserName:            &u.UserName,
		Tags:                u.Tags,
	}
	resp, err := u.Client.CreateUser(&param)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
