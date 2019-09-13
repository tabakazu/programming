package model

type Post struct {
	Title string
	Body  string
}

var (
	post  *Post   = &Post{Title: "TEST 0", Body: "TEST 0"}
	post1 *Post   = &Post{Title: "TEST 1", Body: "TEST 1"}
	post2 *Post   = &Post{Title: "TEST 2", Body: "TEST 2"}
	posts []*Post = []*Post{post, post1, post2}
)

func PostAll() []*Post {
	return posts
}

func PostFind(id int) *Post {
	return posts[id]
}
