package main

import "fmt"

// Duck は、アヒルインターフェイス (アヒル鳴き、アヒル歩き できれば アヒル)
type Duck interface {
	DuckCry()
	DuckWalk()
}

// CryingDuck は、アヒルを鳴かせる関数
func CryingDuck(d Duck) {
	d.DuckCry()
}

// WalkingDuck は、アヒルを歩かせる関数
func WalkingDuck(d Duck) {
	d.DuckWalk()
}

// Person は、人の構造体
type Person struct {
	Name string
}

// DuckCry は、人がアヒル鳴きするメソッド
func (p Person) DuckCry() {
	fmt.Println("アヒルのように鳴く - グエグエ")
}

// DuckWalk は、人がアヒル歩きするメソッド
func (p Person) DuckWalk() {
	fmt.Println("アヒルのように歩く - テクテク")
}

func main() {
	// 人を作成
	person := Person{Name: "Taro"}

	// アヒルを鳴かせる (アヒル鳴きできれば人でもアヒル)
	CryingDuck(person)

	// アヒルを歩かせる (アヒル歩きできれば人でもアヒル)
	WalkingDuck(person)
}
