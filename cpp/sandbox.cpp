#include <bits/stdc++.h>

// 関数
int calculate(int x, int y) {
  return x + y;
}

// クラス
class User {
  private:
    int age;
  public:
    void setage(int new_age);
    int getage();
};

// インスタンスメソッド
void User::setage(int new_age) {
  age=new_age;
}
int User::getage() {
  return age;
}

int main() {
  // 標準入力
  // int a;
  // cin >> a;

  // 標準出力
  std::cout << "Hello World" << std::endl;
    // => Hello World

  ////
  //// 数値
  ////
  int i = 1;
  std::cout << i << std::endl;
    // => 1

  ////
  //// std::string
  ////
  std::string str = "Hello";
  std::cout << str << std::endl;
    // => Hello

  ////
  //// std::vector
  ////
  std::vector<int> vec(3);
  vec.at(0) = 100;
  vec.at(1) = 1;
  vec.at(2) = 10000;
  for (int i = 0; i < vec.size(); i++) {
    std::cout << "vec.at(" << i << ") = " << vec.at(i) << std::endl;
  }
    // => vec.at(0) = 100
    // => vec.at(1) = 1
    // => vec.at(2) = 10000

  // sort : https://cpprefjp.github.io/reference/algorithm/sort.html
  std::sort(vec.begin(), vec.end());
  // reverse : https://cpprefjp.github.io/reference/algorithm/reverse.html
  std::reverse(vec.begin(), vec.end());
  for (int i = 0; i < vec.size(); i++) {
    std::cout << "vec.at(" << i << ") = " << vec.at(i) << std::endl;
  }
    // => vec.at(0) = 10000
    // => vec.at(1) = 100
    // => vec.at(2) = 1


  // メソッド呼び出し
  std::cout << "calculate(1, 2): " << calculate(1, 2) << std::endl;
    // calculate(1, 2): 3

  // インスタンス化
  User u;
  u.setage(20);
  std::cout << "u.getage(): " << u.getage() << std::endl;
    // u.getage(): 20
}
