n = gets.chop.to_i
arr = []

n.times do
  arr.push gets.chop.to_i
end

arr.sort!.reverse!

alice = []
bob = []

n.times do
  alice.push arr.shift unless arr.empty?
  bob.push arr.shift unless arr.empty?
end

puts alice.sum - bob.sum