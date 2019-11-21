class Hoge; end
hoge = Hoge.new

# hoge.fuga
# => `<main>': undefined method `fuga' for #<Hoge:0x0000> (NoMethodError)

class Hoge
  def method_missing(method, *args)
    puts method
  end
end
hoge.send :method_missing, :fuga
# => fuga
