package mock

import "fmt"

// Mock
type RepositoryMock struct{}

func (r *RepositoryMock) Save() string {
	fmt.Println("Not Save")
	s := "Success"
	return s
}
