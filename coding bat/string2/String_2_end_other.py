#Given two strings, return True if either of the strings appears at the very end of the other string, ignoring upper/lower case differences (in other words, the computation should not be "case sensitive"). Note: s.lower() returns the lowercase version of a string.
#end_other('Hiabc', 'abc') → True
#end_other('AbC', 'HiaBc') → True
#end_other('abc', 'abXabc') → True
def end_other(string1,string2):
     num1 = len(string1)
     num2 = len(string2)
     string1 = string1.lower()
     string2 = string2.lower()
     if string1[num2:] == string2[:]:
          return True
     elif string2[num1:] == string1[:]:
          return True
     else:
          return False
print('Enter the two string')
string1 = str(input())
string2 = str(input())
print(str(end_other(string1,string2)))
