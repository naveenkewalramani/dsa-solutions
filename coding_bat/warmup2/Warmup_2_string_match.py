def string_match(string1,string2):
     num1 = len(string1)
     num2 = len(string2)
     count = 0
     if num1 < num2 :
          for i in range(0,num1-1):
               if(string1[i] == string2[i] and string1[i+1] == string2[i+1]):
                    count=count+1
     else :
          for i in range(0,num2-1):
               if(string1[i] == string2[i] and string1[i+1] == string2[i+1]):
                    count = count + 1
     return count
print('Enter the two string')
string1 = str(input())
string2 = str(input())
print(string_match(string1,string2))
