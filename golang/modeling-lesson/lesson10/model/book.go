package model

// 本 (蔵書タイトル)
type Book struct {
	ID    int    // ID
	ISBN  string // ISBN
	Title string // タイトル
}

func NewBook(id int, isbn, title string) *Book {
	return &Book{
		ID:    id,
		ISBN:  isbn,
		Title: title,
	}
}
