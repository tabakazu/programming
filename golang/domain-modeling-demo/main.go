package main

import (
	"fmt"

	"github.com/tabakazu/domain-modeling-demo/domain/model/entity"
	"github.com/tabakazu/domain-modeling-demo/domain/model/valueobject"
	"github.com/tabakazu/domain-modeling-demo/infrastructure/persistence/datastore"
)

func main() {
	// 顧客
	customerRepo := datastore.NewCustomerRepository()

	customerName := valueobject.CustomerName{"New Customer"}
	address := valueobject.Address{Street: "Minato-ku 1-1-1", City: "Tokyo", Country: "Japan"}
	newCustomer := entity.Customer{Name: customerName, Address: address}
	savedCustomer, _ := customerRepo.Save(newCustomer)
	fmt.Println(savedCustomer)

	customers, _ := customerRepo.FindAll()
	fmt.Println(customers)

	// 投稿
	postRepo := datastore.NewPostRepository()

	postName := valueobject.PostName{Value: "New Post"}
	cusID := valueobject.CustomerID{Value: savedCustomer.ID}
	newPost := entity.Post{CustomerID: cusID, Name: postName}
	savedPost, _ := postRepo.Save(newPost)
	fmt.Println(savedPost)

	posts, _ := postRepo.FindAll()
	fmt.Println(posts)

	post1, _ := postRepo.FindByName(postName)
	fmt.Println(post1)
}
