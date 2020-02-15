i, j = gets.split(' ').map(&:to_i)
puts (i * j) % 2 != 1 ? 'Even' : 'Odd'
