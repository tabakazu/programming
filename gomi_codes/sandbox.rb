i = 1
puts i.class, i

f = 0.1
puts f.class, f

s = 'Hello'
puts s.class, s

arr = [1, 2, 3]
puts arr.class, arr

hash = { a: 1, b: 2 }
puts hash.class, hash

Struct.new(:foo, :bar).new('Foo', 'Bar')