s = gets.chop
words = %w[dream dreamer erase eraser]
flag = false

words.map do | word |
  size = word.size 
  if s[-size, size] == word
    s = s[0..-size - 1]
    flag = true
    break
  end
end

puts flag ? 'YES' : 'NO'
