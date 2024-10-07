#You and your date are trying to get a table at a restaurant. The parameter "you" is the stylishness of your clothes, in the range 0..10, and "date" is the stylishness of your date's clothes. The result getting the table is encoded as an int value with 0=no, 1=maybe, 2=yes. If either of you is very stylish, 8 or more, then the result is 2 (yes). With the exception that if either of you has style of 2 or less, then the result is 0 (no). Otherwise the result is 1 (maybe).
#date_fashion(5, 10) → 2
#date_fashion(5, 2) → 0
#date_fashion(5, 5) → 1
def date_fashion(a,b):
     if a > 8 or b > 8:
          return 2
     elif a <=2 or b<=2 :
          return 0
     else:
          return 1
print('Enter the two number')
a = int(input())
b = int(input())
print(date_fashion(a,b))
