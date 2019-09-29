package main

import "github.com/tabakazu/modeling-lesson/lesson03/model"

func main() {
	c := model.NewCustomer(1, "鈴木ハナコ", "hanako@example.jp")
	a := model.NewAddress(1, c.ID, "自宅", "中央区銀座1-1")
	c.Addresses = append(c.Addresses, a)
}
