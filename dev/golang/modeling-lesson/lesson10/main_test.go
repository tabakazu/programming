package main_test

import (
	"testing"

	"github.com/tabakazu/modeling-lesson/lesson10/model"
)

type bookTest struct {
	ID    int
	ISBN  string
	Title string
}

type collectionTest struct {
	ID     int
	BookID int
}

var bookTests = []bookTest{
	{ID: 1, ISBN: "12345", Title: "ぼっちゃん"},
	{ID: 2, ISBN: "24680", Title: "ハリーポッター"},
}

var collectionTests = []collectionTest{
	{ID: 1, BookID: 1},
	{ID: 2, BookID: 1},
}

func TestMain(t *testing.T) {
	for i := range bookTests {
		test := &bookTests[i]
		book := model.NewBook(test.ID, test.ISBN, test.Title)

		if book.ID != test.ID || book.ISBN != test.ISBN || book.Title != test.Title {
			t.Errorf("book.ID, book.ISBN, book.Title = %v, %v, %v want %v, %v, %v", book.ID, book.ISBN, book.Title, test.ID, test.ISBN, test.Title)
		}
	}

	for i := range collectionTests {
		test := &collectionTests[i]
		collection := model.NewCollection(test.ID, test.BookID)

		if collection.ID != test.ID || collection.BookID != test.BookID {
			t.Errorf("collection.ID, collection.BookID = %v, %v want %v, %v", collection.ID, collection.BookID, test.ID, test.BookID)
		}
	}
}
