# メソッドの定義場所をソースコードファイル名と行番号の配列で返す
# https://docs.ruby-lang.org/ja/latest/method/Method/i/source_location.html

module Hoge
  def hoge
    p :hoge_hoge
  end
end

class Fuga
  include Hoge
  def fuga
    p :fuga_fuga
  end
end

class Piyo < Fuga
  def piyo
    p :piyo_piyo
  end
end

pp Piyo.new.method(:hoge).source_location
pp Piyo.new.method(:fuga).source_location
pp Piyo.new.method(:piyo).source_location
