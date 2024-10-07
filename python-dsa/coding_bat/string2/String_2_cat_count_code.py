#Return the number of times that the string "code" appears anywhere in the given string, except we'll accept any letter for the 'd', so "cope" and "cooe" count.
#count_code('aaacodebbb') → 1
#count_code('codexxcode') → 2
#count_code('cozexxcope') → 2
def count_code(string):
     count = 0
     length = len(string)
     for i in range(0,length-3):
          if string[i] == 'c' and string[i+1] == 'o' and string[i+3] == 'e':
               count = count + 1
     return count
print('Enter the string')
string = str(input())
print(str(count_code(string)))
