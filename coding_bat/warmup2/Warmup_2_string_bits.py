#Given a string, return a new string made of every other char starting
#with the first, so "Hello" yields "Hlo".
#string_bits('Hello') → 'Hlo'
#string_bits('Hi') → 'H'
#string_bits('Heeololeo') → 'Hello'

def string_bits(string):
     return string[::2]
print('Enter the string')
string = str(input())
print(str(string_bits(string)))
