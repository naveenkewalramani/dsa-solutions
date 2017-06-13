#Given a string, return the string made of its first two chars, so the String "Hello" yields "He".
#If the string is shorter than length 2, return whatever there is,
#so "X" yields "X", and the empty string "" yields the empty string "".
#first_two('Hello') → 'He'
#first_two('abcdefg') → 'ab'
#first_two('ab') → 'ab'

def first_two(string):
     num = len(string)
     if num < 2:
          return string
     else:
          return string[num-2:]
print('Enter the input string')
string = str(input())
print(str(first_two(string)))
