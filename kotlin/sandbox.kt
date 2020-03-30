//
// クラス
//
class Person(val firstName: String, val lastName: String) {
  fun fullName(): String {
    return "${this.firstName} ${this.lastName}"
  }
}

//
// インターフェース
//
interface IPersonRepository {
  fun findByFirstName(firstName: String): Person
}
// インターフェースの実装
class PersonRepository : IPersonRepository {
  override fun findByFirstName(firstName: String): Person {
    val person = Person(firstName, "Smith")
    return person
  }
}
// インターフェースの実装を呼び出すクラス
class PersonUsecase(val repo: IPersonRepository) {
  fun getPersonByFirstName(firstName: String): Person {
    val person = repo.findByFirstName(firstName)
    return person
  }
}

//
// ジェネリクス
//
interface IGenericsRepository<T> {
  fun first(): T
}
// インターフェースの実装
class PersonGenericsRepository : IGenericsRepository<Person> {
  override fun first(): Person {
    val person = Person("1st", "Person")
    return person
  }
}
// インターフェースの実装を呼び出すクラス
class PersonGenericsUsecase(val repo: IGenericsRepository<Person>) {
  fun getFirstPerson(): Person {
    val person = repo.first()
    return person
  }
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

  // if 式
  val flag: Boolean = true
  if (flag)
    println("flag is true")
  else 
    println("flag is false")

  // クラスのインスタンス生成
  val person = Person("Joe", "Smith")
  println("${person} is ${person.javaClass.name}")
  println("person.firstName is ${person.firstName}")
  println("person.fullName is ${person.fullName()}")

  // インターフェースの実装 & ダックタイピング
  val repo = PersonRepository()
  val usecase = PersonUsecase(repo)
  val person1 = usecase.getPersonByFirstName("Alice")
  println(person1.firstName)

  // ジェネリクス
  val repo2 = PersonGenericsRepository()
  val usecase2 = PersonGenericsUsecase(repo2)
  val person2 = usecase2.getFirstPerson()
  println(person2.fullName())
}
