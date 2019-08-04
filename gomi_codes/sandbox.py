### 型
i = 1
print(i, type(i))
# 出力 : 1 <class 'int'>

s = 'Hello'
print(s, type(s))
# 出力 : Hello <class 'str'>

b = True
print(b, type(b))
# 出力 : True <class 'bool'>

list1 = [1, 'a', True]
print(list1, type(list1))
# 出力 : [1, 'a', True] <class 'list'>

tuple1 = ('Tokyo', 'Osaka')
print(tuple1, type(tuple1))
# 出力 : ('Tokyo', 'Osaka') <class 'tuple'>

dict1 = { 'name':'Taro', 'age':1, 'age':2 }
print(dict1, type(dict1))
# 出力 : {'name': 'Taro', 'age': 2} <class 'dict'>

set1 = { 'Apple', 'Orange', 1, 2, 2 }
print(set1, type(set1))
# 出力 : {1, 2, 'Orange', 'Apple'} <class 'set'>


### 条件分岐 if
x = 1
if x == 0:
  print('x is zero')
elif x == 1:
  print('x is one')
else:
  print('x is other')
# 出力 : x is one


### 繰り返し for
words = ['google', 'apple', 'facebook', 'amazon']
for w in words:
  print(w, end=' ')
# 出力 : google apple facebook amazon

### 繰り返し for + range
for i in range(5):
  print(i, end=' ')
# 出力 : 0 1 2 3 4


### 関数
def greet(name):
  print('Hi!', name)

greet('Alice')
print(type(greet))
# 出力 : Hi! Alice
# 出力 : <class 'function'>

### 関数 キーワード引数
def add_options(id, count, action = 'normal'):
  print(action, count, id)

add_options(count = 2, id = 20)
# add_options(id = 100) これは引数不足でエラー
add_options(action = 'error', count = 1, id = 5)
# 出力 : normal 2 20
# 出力 : error 1 5


### リスト
people = ['Alice', 'Bob']
print(people) # 出力例 : ['Alice', 'Bob']
people.append('Chris')
print(people) # 出力例 : ['Alice', 'Bob', 'Chris']
people.remove('Bob')
print(people) # 出力例 : ['Alice', 'Chris']
people.reverse()
print(people) # 出力例 : ['Chris', 'Alice']
people.insert(1, 'Dave')
print(people) # 出力例 : ['Chris', 'Dave', 'Alice']
people.sort()
print(people) # 出力例 : ['Alice', 'Chris', 'Dave']

for i, v in enumerate(people):
  print(i, v)
# 出力: 0 Alice
# 出力: 1 Chris
# 出力: 2 Dave

### 辞書型 (dictionary)
person = { 'name':'Alice', 'age':20 }
for k, v in person.items():
  print(k, v)
# 出力: name Alice
# 出力: age 20
