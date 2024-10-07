#The parameter weekday is True if it is a weekday, and the parameter vacation is True if we are on vacation.
#We sleep in if it is not a weekday or we're on vacation.
#Return True if we sleep in.
#sleep_in(False, False) → True
#sleep_in(True, False) → False
#sleep_in(False, True) → True

def sleep_in(input1 ,input2):
    if(input1 == False):
        return True
    elif(input1 == True):
        return False
while True:
    print("Enter whether it is a weekday or not ")
    input1 =bool(input())
    if(input1 != False and input1 != True):
        print("Enter the True or False")
        continue
    print("Enter whether it is a vacation or not")
    input2 = bool(input())
    if(input2 != False and input2 != True):
       print("Enter the True or False")
       continue
    print(sleep_in(input1,input2))
