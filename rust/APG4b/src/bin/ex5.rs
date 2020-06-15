// $ cargo run --bin ex5

fn main() {
    let (a, b) = {
        let mut s = String::new();
        std::io::stdin().read_line(&mut s).unwrap();
        let mut iter = s.split_whitespace().map(|i| i.parse::<i32>().unwrap());
        (iter.next().unwrap(), iter.next().unwrap())
    };

    println!("{}", a + b);
}
