console.log('Hello, World')
// => Hello, World
 
const i = 1
console.log(typeof(i), i)
// => number 1

const str = 'Hello'
console.log(typeof(str), str)
// => string Hello

// 配列
const obj1 = [1, 2, 3]
console.log(typeof(obj1), obj1)
// => object [ 1, 2, 3 ]

// 連想配列
const obj2 = { a: 1, b: 2 }
console.log(typeof(obj2), obj2)
// => object { a: 1, b: 2 }


class Person {
  constructor(firstName, lastName, age) {
    this.firstName = firstName
    this.lastName = lastName
    this.age = age
  }

  fullName() {
    return this.firstName + ' ' + this.lastName
  }
}
const person = new Person('taro', 'yamada', 20)
console.log(typeof(person), person)
// => object Person { firstName: 'taro', lastName: 'yamada', age: 20 }
console.log(person.firstName, person.lastName, person.age, person.fullName())
// => taro yamada 20 taro yamada

class User {
  constructor(id) {
    this.id = id
  }
  static find(id) {
    const user = new this(id)
    return user
  }
}
const user = User.find(1)
console.log(typeof(user), user)
// => object User { id: 1 }
