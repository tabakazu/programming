extern crate diesel_demo;
extern crate diesel;

use self::diesel_demo::*;
use self::models::*;
use self::diesel::prelude::*;

fn main() {
    use diesel_demo::schema::items::dsl::*;

    let connection = establish_connection();
    let results = items.select((id, title, body))
        .limit(5)
        .load::<Item>(&connection)
        .expect("Error loading items");
    
    println!("Displaying {} items", results.len());
    for item in results {
        println!("-----------");
        println!("title: {}", item.title);
        println!("body: {}", item.body);
    }
}
