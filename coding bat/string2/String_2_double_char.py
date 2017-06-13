#Given a string, return a string where for every char in the original, there are two chars.
#double_char('The') → 'TThhee'
#double_char('AAbb') → 'AAAAbbbb'
#double_char('Hi-There') → 'HHii--TThheerree'

def double_char(string):
     string2 = ''
     for i in range(0,len(string)):
          string2 = string2 + string[i]*2
     return string2
print('Enter the string')
string = str(input())
print(str(double_char(string)))
