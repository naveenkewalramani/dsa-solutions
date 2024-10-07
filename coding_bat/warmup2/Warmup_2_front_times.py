#Given a string and a non-negative int n, we'll say that the front
#of the string is the first 3 chars, or whatever is
#there if the string is less than length 3.
#Return n copies of the front;
#front_times('Chocolate', 2) → 'ChoCho'
#front_times('Chocolate', 3) → 'ChoChoCho'
#front_times('Abc', 3) → 'AbcAbcAbc'

def front_times(string,num):
     length = len(string)
     if(int(length) < 3):
          return string*int(num)
     else:
          return string[0:3]*int(num)
print('Enter the string')
string  = str(input())
print('Enter the number')
num = int(input())
print(str(front_times(string,num)))
