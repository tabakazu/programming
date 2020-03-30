######
###### 変数の定義
######
a = 1
b = 'a'
c = :xxx
puts "#{a}, #{b} ,#{c}"


######
###### 基本的な型
######
i = 1
str = 'Hello'
sym = :Hello
arr = [1, 'a', :a]
hash = { a: 1, b: 2, c: 3 }
puts "#{i} is #{i.class}"
puts "#{str} is #{str.class}"
puts "#{sym} is #{sym.class}"
puts "#{arr} is #{arr.class}"
puts "#{hash} is #{hash.class}"


######
###### if 式
######
flag = [true, false].sample
if flag
  puts 'flag is true'
else
  puts 'flag is false'
end
# 三項演算子
puts flag ? 'flag is true' : 'flag is false'


######
###### クラス
######
class Person
  # アクセサメソッド
  attr_accessor :first_name, :last_name

  def initialize(first_name, last_name)
    @first_name = first_name
    @last_name = last_name
  end

  def full_name
    first_name + ' ' + last_name
  end
end
# クラスのインスタンス生成
person = Person.new('Joe', 'Smith')
puts "#{person} is #{person.class}"
puts "#{person}#full_name} is #{person.full_name}"


######
###### ダックタイピング
######
class PersonRepository
  def all
    # DB 操作などで取得したとして...
    [Person.new('Taro', 'Yamada')]
  end
end
class PersonResource
  def all
    # API 操作などで取得したとして...
    [Person.new('Hanako', 'Tanaka')]
  end
end
# 呼び出し元クラス
class PersonUsecase
  def initialize(repo)
    @repo = repo
  end
  def get_all
    @repo.all
  end
end
# 喚び出し手から取り換え可能な状況 (多様性)
repo = PersonRepository.new
resouce = PersonResource.new
puts PersonUsecase.new(repo).get_all.map(&:full_name)
puts PersonUsecase.new(resouce).get_all.map(&:full_name)


######
###### 自身をブロックに渡すクラス
######
class Person
  attr_accessor :first_name, :last_name
  def initialize
    yield self if block_given?
  end
end
person = Person.new do |person|
  person.first_name = 'Jiro'
  person.last_name = 'Yamada'
end
puts "#{person}#full_name} is #{person.full_name}" #<Person:xxx>#full_name} is Jiro Yamada
