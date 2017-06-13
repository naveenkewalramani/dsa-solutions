def array_front9(array):
     count = False
     length = len(array)
     if length < 4:
          for i in array:
               if(array[int(i)] == 9):
                    count = True
     else:
          for i in range(0,4):
               if(array[int(i)] == 9):
                    count = True
     return count
array = []
print('Enter the number or -1 to exit')
num = int(input())
while int(num) != -1 :
     array.append(int(num))
     print('Enter the number or -1 to exit')
     num = int(input())
print(str(array_front9(array)))
