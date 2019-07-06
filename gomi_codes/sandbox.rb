# Integer クラス
i = 1
puts i.class, i

# Float クラス
f = 0.1
puts f.class, f

# String クラス
str = 'Hello'
puts str.class, str

# Array クラス
arr = [1, 2, 3]
puts arr.class, arr

# Hash クラス
hash = { a: 1, b: 2 }
puts hash.class, hash


# 構造体クラス
## クラス名あり ※ 第一引数がクラス名、keyword_init が true でキーワード引数有効
person_struct = Struct.new('Person', :name, :age, keyword_init: true)
person1 = person_struct.new(name: 'taro', age: 20)
puts person1.class, person1.name, person1.age

## クラス名なし ※ 定数に入れることで利用可能
Person = Struct.new(:name, :age)
person2 = Person.new('hanako', 40)
puts person2.class, person2.name, person2.age


# 要素を動的に追加・削除できる構造体クラス
require 'ostruct'
o_struct = OpenStruct.new({ name: 'pochi' })
o_struct.age = 10 # 初期値にない要素を追加可能
puts o_struct.class, o_struct.name, o_struct.age