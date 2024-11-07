#Given an array of ints, return True if 6 appears as either the first or last element in the array. The array will be length 1 or more.
#first_last6([1, 2, 6]) → True
#first_last6([6, 1, 2, 3]) → True
#first_last6([13, 6, 1, 2, 3]) → False
def first_last6(array):
     if(int(array[0]) == 6 or int(array[-1]) == 6):
          return True
     else:
          return False
array = []
print('Enter the number or -1 to exit')
num = int(input())
while int(num) != -1:
     array.append(int(num))
     print('Enter the number or -1 to exit')
     num = int(input())
print(str(first_last6(array)))
