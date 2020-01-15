const i: number = 1
console.log(typeof (i), i)

const str: string = 'Hello'
console.log(typeof (str), str)


////
//// クラス
////
class User {
  // プロパティ - 変数名: 型
  firstName: string
  lastName: string

  // コストラクタ（省略可能）
  constructor(firstName: string, lastName: string) {
    this.firstName = firstName
    this.lastName = lastName
  }

  fullName() {
    return this.firstName + ' ' + this.lastName
  }

  static className() {
    return 'userClass'
  }
}

const user = new User('taro', 'yamada')
console.log(typeof(user), user) // object User { firstName: 'taro', lastName: 'yamada' }
console.log(user.firstName, user.lastName, user.fullName()) // taro yamada taro yamada
console.log(User.className()) // userClass


////
//// インターフェース
////
interface IUserRepository {
  find(id: number): string
}

class UserApplicationService {
  repository: IUserRepository

  constructor(userRepository: IUserRepository) {
    this.repository = userRepository
  }

  getUser(id: number) {
    return this.repository.find(id)
  }
}

class UserRepository {
  find(id: number) {
    return `SELECT * FROM users WHERE id = ${id};`
  }
}

const repo = new UserRepository
const service = new UserApplicationService(repo)
console.log(service.getUser(1)) // SELECT * FROM users WHERE id = 1;
