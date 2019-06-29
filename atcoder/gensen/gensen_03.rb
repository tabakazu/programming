n = gets.chop.to_i
arr = []
count = 0 

n.times do
  arr.push gets.chop.to_i
end

while(arr.all? { | x | x % 2 != 1}) do
  arr = arr.map { | x | x / 2 }
  count = count + 1
end

puts count