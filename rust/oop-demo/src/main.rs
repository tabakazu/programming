mod user;

fn main() {
    let user = user::User {
        email: "john_snow@example.com".to_string(),
        family_name: "John".to_string(),
        given_name: "Snow".to_string(),
    };
    println!("Email: {}", user.email);
    println!("Family Name: {}", user.family_name);
    println!("Given Name: {}", user.given_name);
    println!("Full Name: {}", user.full_name());
}
