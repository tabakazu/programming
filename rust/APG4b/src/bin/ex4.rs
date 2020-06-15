// $ cargo run --bin ex4

fn main() {
  // 1 年の秒数
  let seconds: i32 = 365 * 24 * 60 * 60;

  // 1 年は何秒か
  println!("{}", seconds);
  // 2 年は何秒か
  println!("{}", seconds * 2);
  // 5 年は何秒か
  println!("{}", seconds * 5);
  // 10 年は何秒か
  println!("{}", seconds * 10);
}
