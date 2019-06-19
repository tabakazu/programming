def fibo(n)
  n < 2 ? n : fibo(n - 1) + fibo(n - 2)
end

puts fibo(10)
# 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55
