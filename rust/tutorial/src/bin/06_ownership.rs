// cargo run --bin 06_ownership

fn main() {
    // Rust は所有権を使って、ガベージコレクタなしで安全性担保を行うことができる

    // 所有権規則
    // - Rustの各値は、所有者と呼ばれる変数と対応している。
    // - いかなる時も所有者は一つである。
    // - 所有者がスコープから外れたら、値は破棄される。


    let s1 = String::from("hello");
    let s2 = s1; // s1 の値は s2 に束縛
    // println!("{}, world!", s1); // Error. コンパイラはs1が最早有効ではないと考え、参照を無効化
    println!("{}, world!", s2);


    // 整数のようなコンパイル時に既知のサイズを持つ型は、スタック上にすっぽり保持されるので、 実際の値をコピーするのも高速
    let x = 5;
    let y = x;
    println!("x = {}, y = {}", x, y);


    // 所有権と関数
    let s = String::from("hello");
    takes_ownership(s);
    // takes_ownership(s); // Error. ここではもう有効ではない


    // 戻り値の所有権を移動する
    let s1 = String::from("hello");
    let (s2, len) = calculate_length(s1); // s1 を渡して s2 で返却させる
    println!("The length of '{}' is {}.", s2, len);


    // 参照と借用
    let s1 = String::from("hello");
    let len = calculate_length_2(&s1); // 値の所有権をもらう代わりに引数としてオブジェクトへの参照を取る
    println!("The length of '{}' is {}.", s1, len);


    // 可変な参照
    let mut s = String::from("hello");
    change(&mut s);
    println!("{}", s);

    // 参照の規則
    // - 任意のタイミングで、一つの可変参照か不変な参照いくつでものどちらかを行える。
    // - 参照は常に有効でなければならない。
}

fn takes_ownership(some_string: String) {
    println!("{}", some_string);
}

fn calculate_length(s: String) -> (String, usize) {
    let length = s.len();
    (s, length)
}

fn calculate_length_2(s: &String) -> usize {
    s.len()
}

fn change(some_string: &mut String) {
    some_string.push_str(", world !!!!!!!!!");
}
