n = gets.chop.to_i
arr = []

n.times do
  arr.push gets.chop.to_i
end

puts arr.uniq.count