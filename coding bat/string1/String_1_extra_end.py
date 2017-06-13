#Given a string, return a new string made of 3 copies of the last 2 chars
#3of the original string. The string length will be at least 2.
#extra_end('Hello') → 'lololo'
#extra_end('ab') → 'ababab'
#extra_end('Hi') → 'HiHiHi'

def extra_end(string):
     num =len(string)
     string2 = string[num-2:]*3
     return string2
print('Enter the string')
string = str(input())
print(str(extra_end(string)))
