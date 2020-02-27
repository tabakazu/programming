package infra

import "golang.org/x/crypto/bcrypt"

type PasswordDigestGenerator struct{}

func NewPasswordDigestGenerator() *PasswordDigestGenerator {
	return &PasswordDigestGenerator{}
}

func (g *PasswordDigestGenerator) Generate(passwd string) (string, error) {
	passwdByte := []byte(passwd)
	digest, err := bcrypt.GenerateFromPassword(passwdByte, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(digest), nil
}

func GeneratePasswordDigest(passwd string) (string, error) {
	g := NewPasswordDigestGenerator()
	digest, err := g.Generate(passwd)
	if err != nil {
		return "", err
	}
	return digest, nil
}
