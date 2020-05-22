def page(c, t)
  @c = c
  @t = t

  item 1
  item 2 if t >= 2
  item 3 if t >= 3

  item "<" if c >= 6 && c != 6

  for i in (c - 2)..(c + 2) do
    if i <= 3 || i >= t - 2
      next
    end
    item i
  end

  item ">" if t - c >= 5 && t - 5 != c

  item t - 2 if t >= 6
  item t - 1 if t >= 5
  item t if t != 1 && t >= 4

  puts ""
end

def item(n)
  print n == @c ? "_#{n}_" : " #{n} "
end

puts "--------------"

page 1, 7
page 2, 7
page 3, 7
page 4, 7
page 5, 7
page 6, 7
page 7, 7

puts "--------------"

page 1, 8
page 2, 8
page 7, 8
page 8, 8

puts "--------------"

page 1, 14
page 7, 14
page 8, 14
page 14, 14
