#Given an array of ints, return True if the array contains a 2 next to a 2 somewhere.
#has22([1, 2, 2]) → True
#has22([1, 2, 1, 2]) → False
#has22([2, 1, 2]) → False

def has22(array):
     length = len(array)
     for i in range(0,length-1):
          if int(array[i]) == 2 and int(array[i+1]) == 2:
               return True
     return False
array = []
print('Enter the number or -1 to exit')
num = int(input())
while num != -1:
     array.append(num)
     print('Enter the number or -1 to exit')
     num = int(input())
print(str(has22(array)))
