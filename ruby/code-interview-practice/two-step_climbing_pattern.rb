# 1 歩か 2 歩で階段を登れる場合、何通りあるか数える

def calculate(n)
  return 1 if n == 1
  return 2 if n == 2
  2 * n - 3
end

p calculate(1)
p calculate(2)
p calculate(3)
p calculate(4)
p calculate(5)
