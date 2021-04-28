pub struct User {
    pub email: String,
    pub family_name: String,
    pub given_name: String,
}

impl User {
    pub fn full_name(&self) -> String {
        [self.family_name.clone(), self.given_name.clone()].join("")
    }
}

#[test]
fn it_returns_full_name() {
    let user = User {
        email: "john_snow@example.com".to_string(),
        family_name: "John".to_string(),
        given_name: "Snow".to_string(),
    };
    assert_eq!(user.full_name(), "JohnSnow");
}
