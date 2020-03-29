// クラス
class Person(var name: String) {
}

fun main(args: Array<String>) {
  // Hello World
  println("Hello, World!")

  // コメント
  /* コメント */

  // ローカル変数の定義
  val a: Int = 1
  val b = 1  // 型推論
  val c: Int // 初期化
  c = 1      // 代入
  println("${a}, ${b}, ${c}")

  // 基本的な型
  val i: Int = 1
  val str: String = "Hello"
  println("${i} is ${i.javaClass.name}")
  println("${str} is ${str.javaClass.name}")

  // クラスのインスタンス生成
  val person = Person("Joe Smith")
  println("${person} is ${person.javaClass.name}")
  println("person.name is ${person.name}")
}
