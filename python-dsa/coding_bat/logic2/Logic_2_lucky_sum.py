"""Given 3 int values, a b c, return their sum. However, if one of the values is 13 then it does not count towards the sum and values to its right do not count. So for example, if b is 13, then both b and c do not count.
lucky_sum(1, 2, 3) → 6
lucky_sum(1, 2, 13) → 3"""
def lucky_sum(a,b,c):
     if a!=13 and b!=13 and c!=13:
          return a+b+c
     elif a==13 :
          if c == 13:
               return 0
          else:
               return c
     elif b == 13 :
          return a
     elif c==13:
          return a+b
     elif a == 13 and b ==13 and c==13:
          return 0
print('Enter the three numbers')
a = int(input())
b =int(input())
c =int(input())
print(lucky_sum(a,b,c))
