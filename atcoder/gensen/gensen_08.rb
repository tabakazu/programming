n, total = gets.chop.split(' ').map(&:to_i)
x, y, z = -1, -1, -1

0.upto(n) do | i |
  0.upto(n - x) do | j |
    o = n - i - j >= 0 ? n - i - j : 0
    if (10000 * i) + (5000 * j) + (1000 * o) == total
      x, y, z = i, j, o
    end
  end
end

puts "#{x} #{y} #{z}"
