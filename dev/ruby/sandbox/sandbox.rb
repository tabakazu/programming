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


# ブロック
## Proc クラス
proc = Proc.new { | w | puts w }
puts proc.class, proc.call('This is Proc class')

def proc_call_method word, proc
  proc.call word
  proc.call # ブロックの引数が足りなくても大丈夫
end
proc_call_method 'Using proc call method', Proc.new { | w | puts w }

## lambda メソッド例
def lambda_call_method word, _lambda
  _lambda.call word # ブロックの引数が足りないとエラー
end
lambda_call_method 'Using lambda method', lambda { | w | puts w }

## yield 例
def yield_method word
  yield word
end
yield_method('Using yield method') { | w | puts w }

## block 引数例
def block_call_method word, &block
  block.call word
end
block_call_method 'Using block call method' do | w | 
  puts w 
end


# クラス & ダックタイピング
class Item < Struct.new(:id); end
class ItemRepository
  def all; 'SELECT * FROM items'; end
  def find_by_id(id); "SELECT * FROM items WHERE id = #{id}"; end
end
class ItemRepository2
  def all; "Item.all"; end
end
class ItemApplicationService < Struct.new(:repository)
  def get_items; repository.all; end
  def get_item_by_id(id); repository.find_by_id(id); end
end

item_repo = ItemRepository.new
item_service = ItemApplicationService.new(item_repo)
puts item_service.get_items #=> SELECT * FROM it
puts item_service.get_item_by_id(1) #=> SELECT * FROM items WHERE id = 1

item_repo2 = ItemRepository2.new
item_service = ItemApplicationService.new(item_repo2)
puts item_service.get_items #=> Item.all
