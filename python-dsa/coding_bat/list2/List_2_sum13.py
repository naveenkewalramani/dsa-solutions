#Return the sum of the numbers in the array, returning 0 for an empty array. Except the number 13 is very unlucky, so it does not count and numbers that come immediately after a 13 also do not count.
#sum13([1, 2, 2, 1]) → 6
#sum13([1, 1]) → 2
#sum13([1, 2, 2, 1, 13]) → 6
def sum13(array):
     length = len(array)
     total = 0
     for i in range(0,length):
          if int(array[i]) != 13 and int(array[i-1]) != 13:
               total = total + array[i]
     return total
array = []
print('Enter the number or -1 to exit')
num  =int(input())
while num != -1:
     array.append(num)
     print('Enter the number or -1 to exit')
     num = int(input())
print(sum13(array))
