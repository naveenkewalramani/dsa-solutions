"""Given three ints, a b c, return True if one of b or c is "close" (differing from a by at most 1), while the other is "far", differing from both other values by 2 or more. Note: abs(num) computes the absolute value of a number.

close_far(1, 2, 10) → True
close_far(1, 2, 3) → False
close_far(4, 1, 3) → True"""
def close_far(a,b,c):
     if abs(b-a) <= 1:
          if abs(c-a) >=2 and abs(c-b)>=2:
               return True
          else:
               return False
     elif abs(c-a) <=1:
          if abs(b-a) >=2 and abs(b-c) >=2:
               return True
          else:
               return False
     else:
          return False
print('Enter the three numbers')
a = int(input())
b =int(input())
c = int(input())
print(str(close_far(a,b,c)))
