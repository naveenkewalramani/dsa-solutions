def string_splosion(string):
     num = len(string)
     string1 = ''
     for i in range(0,int(num)+1):
          string1 = string1 + string[:i]
     return string1
print('Enter the string')
string = str(input())
print(str(string_splosion(string)))
