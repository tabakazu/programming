package model

// 蔵書
type Collection struct {
	ID     int // ID
	BookID int // 本の ID
	// BoughtAt // 買った日
}

func NewCollection(id int, book_id int) *Collection {
	return &Collection{
		ID:     id,
		BookID: book_id,
	}
}
