// cargo run --bin 04_functions

fn main() {
    my_print(5);
    my_print_2("hoge");
    my_print(one());
    my_print(plus_one(1));
    early_return();
}

fn my_print(i: i32) {
    println!("`my_print({})` print {}", i, i);
}

fn my_print_2(s: &str) {
    println!("`my_print_2(\"{}\")` print {}", s, s);
}

fn one() -> i32 {
    // 返却値にセミコロン(;)不要
    1
}

fn plus_one(x: i32) -> i32 {
    x + 1
}

fn early_return() {
    println!("<----- ここは出力される ----->");
    // 早期リターン
    return ;

    println!("<----- ここは出力されない ----->");
}
