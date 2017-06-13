#Given 2 int values, return True if one is negative and one is positive.
#Except if the parameter "negative" is True,
#then return True only if both are negative.
#pos_neg(1, -1, False) → True
#pos_neg(-1, 1, False) → True
#pos_neg(-4, -5, True) → True


def pos_neg(num1,num2,negative):
     if (int(num1)<0 and int(num2)>=0) or(int(num1)>=0 and int(num2)<0) :
          return True
     elif (int(num1) < 0 and int(num2)<0 and bool(negative) == True):
          return True
     else:
          return False

print(str(pos_neg(1,-1,False)))
print(str(pos_neg(-1,1,False)))
print(str(pos_neg(-4,-5,True)))
