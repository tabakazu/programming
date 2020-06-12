table! {
    items (id) {
        id -> Integer,
        title -> Varchar,
        body -> Text,
        created_at -> Timestamp,
    }
}

table! {
    posts (id) {
        id -> Integer,
        title -> Varchar,
        body -> Text,
        published -> Bool,
    }
}

allow_tables_to_appear_in_same_query!(
    items,
    posts,
);
