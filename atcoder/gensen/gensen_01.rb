a, b = gets.chop.split(' ').map(&:to_i)
if (a * b) % 2 == 1
  puts 'Odd'
elsif (a * b) % 2 == 0
  puts 'Even'
end