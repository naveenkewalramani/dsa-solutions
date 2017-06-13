#Return the sum of the numbers in the array, except ignore sections of numbers starting with a 6 and extending to the next 7 (every 6 will be followed by at least one 7). Return 0 for no numbers.

#sum67([1, 2, 2]) → 5
#sum67([1, 2, 2, 6, 99, 99, 7]) → 5
#sum67([1, 1, 6, 7, 2]) → 4
def sum67(array):
     length = len(array)
     total = 0
     for i in range(0,length):
          if int(array[i]) == 6:
               array[i] = 0
               k = i + 1
               while int(array[k]) != 7 :
                    array[k] = 0
                    k = k +1
     for i in range(0,length):
          if int(array[i]) != 7:
               total = total + int(array[i])
     return total
print('Enter the hnumber or -1 to exit')
array = []
num = int(input())
while num != -1:
     array.append(num)
     print('Enter the number or -1 to exit')
     num = int(input())
print(sum67(array))
