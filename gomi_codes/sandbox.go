package main

import "fmt"

type People struct {
	Name string
	Age  int
	Sex  string
}

// 値レシーバ
func (p People) changeNameValue(name string) {
	p.Name = name
}

// ポインタレシーバ
func (p *People) changeNamePointer(name string) {
	p.Name = name
}

func (p People) getName() {
	fmt.Println(p.Name)
}

type Dog struct{ Name string }

func (d Dog) getName() {
	fmt.Println(d.Name)
}

type Animal interface {
	getName()
}

func greeting(a Animal) {
	a.getName()
}

func main() {
	i := 1 // int
	fmt.Printf("%T - %v\n", i, i)

	f64 := 0.5 // float64
	fmt.Printf("%T - %v\n", f64, f64)

	s := "Hello" // string
	fmt.Printf("%T - %v\n", s, s)

	arr1 := [5]int{1, 2, 3} // array : 要素が決まっている。
	fmt.Printf("%T - %v\n", arr1, arr1)

	arr2 := [5]string{"alice", "bob", "carol"} // array
	fmt.Printf("%T - %v\n", arr2, arr2)

	sli := []int{1, 2, 3} // slice : 要素数が決まってない。
	sli = append(sli, 4)  // 追加できる。
	fmt.Printf("%T - %v\n", sli, sli)

	m := make(map[string]int) // map
	m["one"] = 1              // 代入できる
	m["two"] = 2
	fmt.Printf("%T - %v\n", m, m)

	people1 := People{"Taro", 15, "Male"} // 構造体 初期化 (初期値の省略不可)
	fmt.Printf("%T - %v\n", people1, people1)

	people2 := People{Name: "Hanako", Age: 60} // 構造体 初期化 (フィールドを指定すれば、初期値の省略可)
	people2.Sex = "Female"                     // 途中で代入可
	fmt.Printf("%T - %v\n", people2, people2)

	people3 := People{Name: "Jiro", Age: 25, Sex: "Female"}
	people3.changeNameValue("Saburo")         // 値レシーバ
	fmt.Printf("%T - %v\n", people3, people3) // 変わらない

	people4 := People{Name: "Jiro", Age: 25, Sex: "Female"}
	people4.changeNamePointer("Saburo")       // ポインタレシーバ
	fmt.Printf("%T - %v\n", people4, people4) // 変わる

	peoples := []People{people1, people2, people3, people4}
	for i, people := range peoples { // slice のループ
		fmt.Println(i, people)
	}

	// ポリモーフィズム
	taro := People{Name: "Taro"}
	pochi := Dog{Name: "Pochi"}
	greeting(taro)
	greeting(pochi)
}
