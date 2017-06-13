#Given 2 ints, a and b, return True if one if them is 10 or if their sum is 10.
#makes10(9, 10) → True
#makes10(9, 9) → False
#makes10(1, 9) → True

def makes10(num1,num2):
    if(int(num1) == 10 or int(num2) == 10 or int(num1)+int(num2) == 10):
        return True
    else:
        return False

print("Enter any two integer")
num1 = int(input())
num2 = int(input())
print(str(makes10(num1,num2)))
