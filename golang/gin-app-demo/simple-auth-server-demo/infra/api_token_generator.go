package infra

import "github.com/google/uuid"

type APITokenGenerator struct{}

func NewAPITokenGenerator() *APITokenGenerator {
	return &APITokenGenerator{}
}

func (g *APITokenGenerator) Generate() (string, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	token := u.String()
	return token, nil
}

func GenerateAPIToken() (string, error) {
	g := NewAPITokenGenerator()
	token, err := g.Generate()
	if err != nil {
		return "", err
	}
	return token, nil
}
