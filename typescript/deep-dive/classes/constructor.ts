// コンストラクタ

class User {
  name: string;
  age: number;
  private isAdmin: boolean;
  constructor(name: string, age: number, isAdmin: boolean) {
    this.name = name;
    this.age = age;
    this.isAdmin = isAdmin;
  }
}

var user: User = new User('Taro', 25,  false);
console.log(user);
console.log(user.name);
console.log(user.age);

// 出力
// User { name: 'Taro', age: 25, isAdmin: false }
// Taro
// 25
