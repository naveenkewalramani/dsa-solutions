def array123(array):
     count = False
     for i in range(0,len(array)-3):
          if( int(array[i]) == 1 and int(array[i+1]) == 2 and int(array[i+2]) == 3):
               count = True
               break
     return count
array = []
print('Enter the number and -1 to exit')
num = int(input())
while int(num) != -1:
     array.append(int(num))
     print('Enter the number and -1 to exit')
     num = int(input())
print(str(array123(array)))     
          
