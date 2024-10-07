#Given an array of ints, return the sum of the first 2 elements in the array. If the array length is less than 2, just sum up the elements that exist, returning 0 if the array is length 0.
#sum2([1, 2, 3]) → 3
#sum2([1, 1]) → 2
#sum2([1, 1, 1, 1]) → 2

def sum2(a):
     length = len(a)
     total=0
     if(length == 0):
          return 0
     else:
          for i in range(0,2):
               total = total + int(a[i])
          return total
print(sum2([1,2,3]))
print(sum2([1,1]))
print(sum2([1,1,1,1]))
