#We have a loud talking parrot. The "hour" parameter is the current hour time in the range 0..23We are in trouble if the parrot is talking and the hour is before 7 or after 20.
#Return True if we are in trouble.
#parrot_trouble(True, 6) → True
#parrot_trouble(True, 7) → False
#parrot_trouble(False, 6) → False

def parrot_trouble(input1,input2):
    if(bool(input1) == False):
        return False
    elif(bool(input1) == True):
        if(int(input2) < 7 or int(input2) > 20):
            return True
        else:
            return False

print("Enter wheteher the parrot is speaking")
input1 = bool(input())
print("Enter the hour from 0 to 23 and -1 to exit")
input2 = int(input())
while int(input2) != -1:
    print(str(parrot_trouble(input1,input2)))
    print("Enter wheteher the parrot is speaking")
    input1 = bool(input())
    print("Enter the hour from 0 to 23 and -1 to exit")
    input2 = int(input())
