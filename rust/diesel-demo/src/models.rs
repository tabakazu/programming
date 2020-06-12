#[derive(Queryable)]
pub struct Item {
    pub id: i32,
    pub title: String,
    pub body: String,
}

#[derive(Queryable)]
pub struct Post {
    pub id: i32,
    pub title: String,
    pub body: String,
    pub published: bool,
}
