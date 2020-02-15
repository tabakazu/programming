package model_test

import (
	"testing"

	"github.com/tabakazu/practice20191209/domain/model"
)

type userTest struct {
	Email    string
	Password string
}

var userTests = []userTest{
	{Email: "Alice@example.org", Password: "Pa$$w0rd"},
}

func TestNewUser(t *testing.T) {
	for i := range userTests {
		test := &userTests[i]
		u := model.NewUser(test.Email, test.Password)

		if u.Email != test.Email || u.Password != test.Password {
			t.Errorf("u.Email, u.Password = %v, %v want %v, %v", u.Email, u.Password, test.Email, test.Password)
		}
	}
}

func TestUserAuthenticate(t *testing.T) {
	for i := range userTests {
		test := &userTests[i]
		u, _ := model.UserAuthenticate(test.Email, test.Password)

		if u.PasswordDigest == "" {
			t.Errorf("u.PasswordDigest is present")
		}
	}
}
