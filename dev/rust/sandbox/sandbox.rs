// 関数
fn add_one(x: i32) -> i32 {
    x + 1
}

// 構造体 (Struct)
struct Post {
    title: String,
    body: String
}

// 列挙型 (Enum)
#[derive(Debug)]
enum IpAddr {
    V4(u8, u8, u8, u8),
}

fn main() {
    // 数値
    let i: i32 = 5;
    println!("{}", i); // => 5

    // 文字列
    let s: String = "Hello, Rust!".to_string();
    println!("{}", s); // => Hello, Rust!

    // スコープ
    {
        let s = "Hello, Rust !!!!!";
        println!("{}", s); // => Hello, Rust !!!!!
    }

    // if
    let condition = true;
    if condition {
        println!("{} to be true", condition); // => true to be true
    } else {
        println!("{} to not be true", condition); // => false to not be true
    }

    // 関数
    println!("{}", add_one(1)); // => 2

    // 構造体
    let post = Post {
        title: String::from("Test title"),
        body: String::from("Test body"),
    };
    println!("{}", post.title); // => Test title
    println!("{}", post.body); // => Test body

    // 列挙型
    let v4_addr = IpAddr::V4(192, 168, 16, 1);
    println!("{:?}", v4_addr); // => V4(192, 168, 16, 1)

    // マッチ
    let x = 3;
    match x {
        1 | 2 => println!("one or two"),
        3 => println!("three"), // => three
        _ => println!("anything"),
    }
}
