#Given 2 arrays of ints, a and b, return True if they have the same first element or they have the same last element. Both arrays will be length 1 or more.
#common_end([1, 2, 3], [7, 3]) → True
#common_end([1, 2, 3], [7, 3, 2]) → False
#common_end([1, 2, 3], [1, 3]) → True
def common_end(a,b):
     if int(a[0]) == int(b[0]) or int(a[-1]) == int(b[-1]):
          return True
     else:
          return False
a = []
b = []
print('Enter number for array1 and -1 to exit')
num  = int(input())
while int (num) != -1:
     a.append(int(num))
     print('Enter number for array1 and -1 to exit')
     num  = int(input())
print('Enter number for array2 and -1 to exit')
num  = int(input())
while int (num) != -1:
     b.append(int(num))
     print('Enter number for array2 and -1 to exit')
     num  = int(input())
print(str(common_end(a,b)))
