package main

import "fmt"

type Post struct {
	Name string
}

func HandleData(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Printf("数字だぞ。%t\n", x)
	case string:
		fmt.Printf("文字だぞ。%t\n", x)
	case Post:
		fmt.Printf("Post型だぞ。%t\n", x)
	}
}

func main() {
	i := 5
	HandleData(i)

	s := "Hello"
	HandleData(s)

	post := Post{Name: "Taro"}
	HandleData(post)
}
