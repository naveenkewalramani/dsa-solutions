#Given a string, return a version without the first and last char, so "Hello" yields "ell". The string length will be at least 2.
#without_end('Hello') → 'ell'
#without_end('java') → 'av'
#without_end('coding') → 'odin'

def without_end(string):
     return string[1:len(string)-1]
print('Enter the string')
string = str(input())
print(str(without_end(string)))
