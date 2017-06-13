#Return the number of times that the string "hi" appears anywhere in the given string.
#count_hi('abc hi ho') → 1
#count_hi('ABChi hi') → 2
#count_hi('hihi') → 2
def count_hi(string):
     count = 0
     length = len(string)
     for i in range(0,length-1):
          if(string[i] == 'h' and string[i+1] == 'i'):
              count = count+1
     return count
print('Enter the string')
string = str(input())
print(count_hi(string))
