package main

import "fmt"

// User はユーザの構造体
type User struct {
	Name string
}

// SetName は User.Name を設定するメソッド
func (p *User) SetName(name string) {
	p.Name = name
	return
}

func main() {
	var p1 *int
	i := 1
	p1 = &i
	fmt.Printf("i    is %v\n*p1  is %v\n", i, *p1)
	*p1 = 2
	fmt.Printf("i    is %v\n*p1  is %v\n", i, *p1)

	var p2 *User
	usr1 := User{"Taro"}
	p2 = &usr1
	fmt.Printf("usr1 is %v\n*p2  is %v\n", usr1, *p2)
	(*p2).Name = "Jiro"
	fmt.Printf("usr1 is %v\n*p2  is %v\n", usr1, *p2)

	p3 := &User{"Hanako"}
	fmt.Printf("*p3  is %v\n", *p3)
	(*p3).Name = "Sakiko"
	fmt.Printf("*p3  is %v\n", *p3)

	usr2 := User{"Alice"}
	fmt.Printf("usr2 is %v\n", usr2)
	usr2.SetName("Bob")
	fmt.Printf("usr2 is %v\n", usr2)
}
