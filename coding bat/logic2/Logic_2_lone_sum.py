"""We want to make a row of bricks that is goal inches long. We have a number of small bricks (1 inch each) and big bricks (5 inches each). Return True if it is possible to make the goal by choosing from the given bricks. This is a little harder than it looks and can be done without any loops. See also: Introduction to MakeBricks
make_bricks(3, 1, 8) → True
make_bricks(3, 1, 9) → False
make_bricks(3, 2, 10) → True"""
def lone_sum(a,b,c):
     if a == b and b == c:
          return 0
     elif a == b :
          return c
     elif b==c :
          return a
     elif a == c:
          return b
     else:
          return a+b+c
print('Enter the three numbers')
a = int(input())
b = int(input())
c = int(input())
print(lone_sum(a,b,c))
