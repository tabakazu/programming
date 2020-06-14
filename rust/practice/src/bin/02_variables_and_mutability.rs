// cargo run --bin 02_variables_and_mutability

fn main() {
    // 変数。デフォルトはイミュータブル(不変)なので再代入不可
    let x = 5;
    println!("x is: {}", x);

    // ミュータブル(可変)な変数を宣言
    let mut y = 5;
    println!("y is: {}", y);
    y = 999;
    println!("y is: {}", y);

    // シャドーイング。前に定義した変数と同じ名前の変数を新しく宣言
    let z = 5;
    println!("z is: {}", z);
    let z = 999;
    println!("z is: {}", z);
}
