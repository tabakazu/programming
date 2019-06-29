n = gets.chop.to_i
min = gets.chop.to_i
max = gets.chop.to_i
total = 0

0.upto(n) do | i |
  sum = i.to_s.split('').map(&:to_i).sum
  total = total + i if min <= sum && max >= sum
end

puts total