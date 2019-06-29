c500_count = gets.chop.to_i
c100_count = gets.chop.to_i
c50_count = gets.chop.to_i
total = gets.chop.to_i
count = 0

0.upto(c500_count) do | i |
  0.upto(c100_count) do | x |
    0.upto(c50_count) do | y |
      _total = (500 * i) + (100 * x) + (50 * y)
      count = count + 1 if total == _total
    end
  end
end

puts count