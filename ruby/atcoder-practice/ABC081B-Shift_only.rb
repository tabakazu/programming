ABC081B-Shift_only.rb
gets.chop.to_i
s = gets.chop.split(' ').map(&:to_i)
@min = 1000000000
def f(cnt, v)
  v % 2 == 1 ? cnt : f(cnt + 1, v / 2)
end
s.each do |n|
  cnt = f(0, n)
  @min = cnt if cnt < @min
end
puts @min
