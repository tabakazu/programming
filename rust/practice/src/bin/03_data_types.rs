// cargo run --bin 03_data_types

use std::any::type_name;

fn type_of<T>(_: T) -> &'static str {
    type_name::<T>()
}

fn main() {
    //
    // Rustは静的型付き言語で、使用したい型を推論してくれる。
    //

    // 整数型
    let i = 127;
    println!("{} is {} type", i, type_of(i));

    // 浮動小数点型
    let j = 3.2;
    println!("{} is {} type", j, type_of(j));

    // 論理値型 boolean
    let t = true;
    let f = false;
    println!("{} is {} type", t, type_of(t));
    println!("{} is {} type", f, type_of(f));

    // 文字型
    let s = "hogehoge";
    println!("{} is {} type", s, type_of(s));

    // 複合型
    // x.0, x.1 ... のように取得できる
    let tup: (i32, f64, &str) = (500, 6.4, "hoge");
    // let tup = (500, 6.4, "hoge");
    println!("{} is {} type", tup.0, type_of(tup.0));
    println!("{} is {} type", tup.1, type_of(tup.1));
    println!("{} is {} type", tup.2, type_of(tup.2));

    // 配列型
    let arr: [i32; 5] = [1, 2, 3, 4, 5];
    // let arr = [1, 2, 3, 4, 5];
    println!("{} is {} type", arr[0], type_of(arr[0]));
    println!("{} is {} type", arr[1], type_of(arr[1]));

    //
    // 型推論ができないパターン
    //
    // 複数の型が可能な場合には、型注釈をつけなければいけません。
    let guess: u32 = "42".parse().expect("Not a number!");
    // let guess = "42".parse().expect("Not a number!"); // 型注釈がないのでエラー
    println!("guess is: {}", guess);
}
