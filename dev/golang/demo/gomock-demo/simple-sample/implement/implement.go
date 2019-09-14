package implement

import "fmt"

// implement
type RepositoryImpl struct{}

func (r *RepositoryImpl) Save() string {
	fmt.Println("Save")
	s := "Success"
	return s
}
