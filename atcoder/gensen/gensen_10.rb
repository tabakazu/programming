n = gets.chop.to_i
arr = []
arr.push [0, 0, 0]
flag = true

n.times do
  arr.push gets.chop.split(' ').map(&:to_i)
end

n.times do | i |
  dt = arr[i + 1][0] - arr[i][0]
  dist = (arr[i + 1][1] - arr[i][1]).abs + (arr[i + 1][2] - arr[i][2]).abs
  flag = false if dt < dist
  flag = false if dist % 2 != dt % 2
end

puts flag ? 'Yes' : 'No'