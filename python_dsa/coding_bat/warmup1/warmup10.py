#Given a string, return a new string where the first and last chars have
#been exchanged.
#front_back('code') → 'eodc'
#front_back('a') → 'a'
#front_back('ab') → 'ba'

def front_back(string):
     return string[::-1]
string1 = input()
string2 = front_back(string1)
print(string2)
