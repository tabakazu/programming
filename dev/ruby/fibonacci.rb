def rec(f2, f1, n)
  n == 1 ? f1 : rec(f1, f2 + f1, n - 1)
end

def fibo(n)
  n < 2 ? n : rec(0, 1, n)
end

puts fibo(1000)