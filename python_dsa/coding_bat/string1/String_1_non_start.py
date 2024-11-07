#Given 2 strings, return their concatenation, except omit the
#first char of each. The strings will be at least length 1.
#non_start('Hello', 'There') → 'ellohere'
#non_start('java', 'code') → 'avaode'
#non_start('shotl', 'java') → 'hotlava
def non_start(string1,string2):
     return string1[1:] + string2[1:]
print('Enter the two string')
string1 = str(input())
string2 = str(input())
print(str(non_start(string1,string2)))

