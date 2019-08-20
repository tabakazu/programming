package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/iam/iamiface"
	"github.com/jessevdk/go-flags"
)

type options struct {
	Username            string `short:"u" long:"username" description:"username"`
	Path                string `short:"p" long:"path" description:"path"`
	PermissionsBoundary string `short:"b" long:"permissionsboundary" description:"permissions boundary"`
	TagName             string `short:"t" long:"tag" description:"tag"`
}

func main() {
	var opts options
	if _, err := flags.Parse(&opts); err != nil {
		fmt.Println(err)
		return
	}

	tagKey := "Name"
	iamTag := iam.Tag{
		Key:   &tagKey,
		Value: &opts.TagName,
	}
	iamTags := []*iam.Tag{&iamTag}

	user := &User{
		Client:              iam.New(session.New()),
		Path:                opts.Path,
		PermissionsBoundary: opts.PermissionsBoundary,
		UserName:            opts.Username,
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
