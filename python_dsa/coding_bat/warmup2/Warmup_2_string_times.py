#Given a string and a non-negative int n,
#return a larger string that is n copies of the original string.
#string_times('Hi', 2) → 'HiHi'
#string_times('Hi', 3) → 'HiHiHi'
#string_times('Hi', 1) → 'Hi'


def string_times(string,num):
     return string*int(num)

print('Enter a string')
string = str(input())
print('Enter a non negative integer')
num = int(input())
print(str(string_times(string,num)))
