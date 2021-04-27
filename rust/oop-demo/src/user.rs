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
