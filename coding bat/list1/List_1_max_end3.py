def max_end3(array):
     if(int(array[0]) > int(array[2])):
          b = [int(array[0]),int(array[0]),int(array[0])]
     else:
          b = [int(array[2]),int(array[2]),int(array[2])]
     return b
a = [1,2,3]
print(max_end3(a))
a = [11,5,9]
print(max_end3(a))
a = [2,11,3]
print(max_end3(a))
