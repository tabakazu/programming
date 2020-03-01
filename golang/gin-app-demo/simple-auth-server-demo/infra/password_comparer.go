package infra

import "golang.org/x/crypto/bcrypt"

type PasswordComparer struct{}

func NewPasswordComparer() *PasswordComparer {
	return &PasswordComparer{}
}

func (c *PasswordComparer) Compare(passwd, digest string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(digest), []byte(passwd)); err != nil {
		return err
	}
	return nil
}

func ComparePassword(passwd, digest string) error {
	c := NewPasswordComparer()
	if err := c.Compare(passwd, digest); err != nil {
		return err
	}
	return nil
}
