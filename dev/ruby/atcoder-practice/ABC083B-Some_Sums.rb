n, a, b = gets.chop.split(' ').map(&:to_i)
@total = 0

1.upto(n) do |i|
  n1 = i % 10
  n2 = i % 100 / 10
  n3 = i % 1000 / 100
  n4 = i % 10000 / 1000
  n5 = i / 10000
  t = n1 + n2 + n3 + n4 + n5
  @total += i if a <= t && t <= b
end

puts @total
