def cat_dog(string):
     count1 = 0
     count2 = 0
     length = len(string)
     for i in range(0,length-2):
          if string[i:i+3] == 'cat':
               count1 = count1 + 1
          elif string[i:i+3] == 'dog':
               count2 = count2 + 1
     if count1 == count2:
          return True
     else:
          return False
print('Enter the string')
string = str(input())
print(str(cat_dog(string)))
