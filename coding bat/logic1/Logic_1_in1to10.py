#Given a number n, return True if n is in the range 1..10, inclusive. Unless outside_mode is True, in which case return True if the number is less or equal to 1, or greater or equal to 10.
#in1to10(5, False) → True
#in1to10(11, False) → False
#in1to10(11, True) → True
def in1to10(num,logic):
     if bool(logic) == False:
          if num>=1 and num<=10:
               return True
          else:
               return False
     else:
          if num<=1 and num>=10:
               return True
          else:
               return False
print('Enter the number and than the logic')
num  = int(input())
logic = bool(input())
print(str(in1to10(num,logic)))
