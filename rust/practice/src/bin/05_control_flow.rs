// cargo run --bin 05_control_flow

fn check_same(actual: i32, expected: i32) {
    // if 式
    if actual == expected {
        println!("{}と{}は同じ", actual, expected); // 条件は真でした
    } else {
        println!("{}と{}は違う", actual, expected); // 条件は偽でした
    }
}

fn check_same_2(actual: i32, expected: i32) {
    // else if
    if actual == expected {
        println!("{}と{}は同じ", actual, expected);
    } else if actual < expected {
        println!("{}は{}より小さいです", actual, expected);
    } else {
        println!("{}は{}より大きいです", actual, expected);
    }
}

fn main() {
    check_same(5, 5);
    check_same(5, 10);

    check_same_2(5, 5);
    check_same_2(5, 10);
    check_same_2(5, 0);

    // if の上限によって変数の中身を変更
    let b: bool = if true {
        true
    } else {
        false
    };
    println!("{}", b);
}
