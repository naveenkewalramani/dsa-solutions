def array_count9(array):
     count = 0;
     for i in range(0,len(array)):
          if int(array[int(i)]) == 9:
               count = count + 1
     return count
array = []
print('Enter the number and -1 to exit')
num = int(input())
while int(num) != -1:
     array.append(int(num))
     print('Enter the number and -1 to exit')
     num = int(input())
print(str(array_count9(array)))     
