package main

import (
	"fmt"

	"github.com/tabakazu/gomock-demo-simple-sample/implement"
	"github.com/tabakazu/gomock-demo-simple-sample/mock"
)

// モック(Mock)と実装(Implement)に共通するインターフェイス
type Repository interface {
	Save() string
}

func main() {
	// Repositoryを呼び出すクライアント
	type RepositoryClient struct {
		repo Repository
	}

	// Using Mock
	c := &RepositoryClient{
		repo: &mock.RepositoryMock{},
	}
	s := c.repo.Save() // 実装と同じ挙動を返すが、破壊的な処理は行われない
	fmt.Printf("Success == %v\n", s)

	// Using Implement
	c.repo = &implement.RepositoryImpl{}
	s = c.repo.Save() // 実装通りの挙動が起きる
	fmt.Printf("Success == %v\n", s)
}
