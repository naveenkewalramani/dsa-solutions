#Return True if the given string contains an appearance of "xyz" where the xyz is not directly preceeded by a period (.). So "xxyz" counts but "x.xyz" does not.
#xyz_there('abcxyz') → True
#xyz_there('abc.xyz') → False
#xyz_there('xyz.abc') → True
def xyz_there(string):
     length = len(string)
     for i in range(0,length):
          if string[i-1] != '.' and string[i:i+3] == 'xyz' :
               return True
          else :
               continue
     return False
print('Enter the string')
string = str(input())
print(str(xyz_there(string)))
