package main

import "fmt"

// User ユーザ構造体
type User struct {
	Name string
}

// Greet ユーザが挨拶するポインタレシーバメソッド
func (u *User) Greet() string {
	return fmt.Sprintf("My name is %v.", u.Name)
}

// Admin 管理者ユーザ構造体
type Admin struct {
	User
}

func main() {
	// 初期化
	admin := &Admin{
		User{
			Name: "Taro",
		},
	}

	// admin.Name で User の Name を呼べる
	fmt.Printf("My name is %v\n", admin.Name)

	// admin.Greet() で User の Greet() が呼べる
	fmt.Printf("%v\n", admin.Greet())

	// admin.Name で User の Name を書き換えれる
	admin.Name = "Hanako"
	fmt.Printf("My name is %v\n", admin.Name)

	// User を初期値にしなくても大丈夫
	admin2 := &Admin{}
	admin2.Name = "Jiro"
	fmt.Printf("My name is %v\n", admin2.Name)
}
