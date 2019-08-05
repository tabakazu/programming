# コンポジション サンプル

class FullCourse
  attr_reader :foods

  def initialize foods
    @foods = foods
  end

  def zensais
    foods.zensais
  end

  def mains
    foods.mains
  end
end

class Foods
  attr_reader :foods

  def initialize foods
    @foods = foods
  end

  def zensais
    foods.select { | food | food.type == 'zensai' }
  end

  def mains
    foods.select { | food | food.type == 'main' }
  end
end

class Food
  attr_reader :name, :type

  def initialize args
    @name = args[:name]
    @type = args[:type]
  end
end

# ---------------------- #

# Food を複数作成
salad1 = Food.new({ name: 'シーザサラダ', type: 'zensai' })
salad2 = Food.new({ name: '和風サラダ', type: 'zensai' })
steak = Food.new({ name: 'ステーキ', type: 'main' })
# コース用の Food リストを作成
course_foods = Foods.new [salad1, salad2, steak]
# コースを作成
course = FullCourse.new course_foods

# コースの前菜を表示
course.zensais.each do | zensai |
  puts zensai.name
end

# コースの主食を表示
course.mains.each do | _main |
  puts _main.name
end

