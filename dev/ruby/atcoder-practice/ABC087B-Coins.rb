A = gets.chop.to_i
B = gets.chop.to_i
C = gets.chop.to_i
X = gets.chop.to_i
@cnt = 0
0.upto(A) do |i|
  0.upto(B) do |j|
    0.upto(C) do |k|
      total = 500*i + 100*j + 50*k
      @cnt += 1 if X == total
    end
  end
end
puts @cnt
