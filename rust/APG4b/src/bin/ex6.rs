// $ cargo run --bin ex6

fn main() {
    let (a, op, b) = {
        let mut s = String::new();
        std::io::stdin().read_line(&mut s).unwrap();
        let mut iter = s.split_whitespace();
        let a: i32 = iter.next().unwrap().parse().unwrap();
        let op: String = iter.next().unwrap().parse().unwrap();
        let b: i32 = iter.next().unwrap().parse().unwrap();
        (a, op, b)
    };

    if op == "+" {
        println!("{}", a + b);
    } else if op == "-" {
        println!("{}", a - b);
    } else if op == "*" {
        println!("{}", a * b);
    } else if op == "/" {
        if b == 0 {
            println!("error");
        } else {
            println!("{}", a / b);
        }
    } else {
        println!("error");
    }
}
