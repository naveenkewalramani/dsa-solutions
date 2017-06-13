#Given an array of ints, return True if the array is length 1 or more,
#and the first element and the last element are equal.
#same_first_last([1, 2, 3]) → False
#same_first_last([1, 2, 3, 1]) → True
#same_first_last([1, 2, 1]) → True

def same_first_last(array):
     length = len(array)
     if length >= 1 and int(array[0]) == int(array[-1]):
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

print(str(same_first_last(array)))
