
#We have two monkeys, a and b, and the parameters a_smile and b_smile indicate if each is smiling.
#We are in trouble if they are both smiling or if neither of them is smiling.
#Return True if we are in trouble.

#monkey_trouble(True, True) → True
#monkey_trouble(False, False) → True
#monkey_trouble(True, False) → False

def monkey_trouble(smile1,smile2):
    if(smile1 == True and smile2 == True):
        return True
    elif(smile1 == False and smile2 == True):
        return False
    elif(smile1 == True and smile2 == False):
        return False
    elif(smile1 == False and smile2 == False):
        return True
print(str(monkey_trouble(True,False)))
print(str(monkey_trouble(True,True)))
print(str(monkey_trouble(False,True)))
print(str(monkey_trouble(False,False)))
