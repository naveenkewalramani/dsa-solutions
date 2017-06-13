#Given an int n, return True if it is within 10 of 100 or 200.
#Note: abs(num) computes the absolute value of a number.
#near_hundred(93) → True
#near_hundred(90) → True
#near_hundred(89) → False

def near_hundred(num):
     if(abs(int(num) - 100)) < 10:
          return True
     elif(abs(int(num) - 200)) < 10:
          return True
     else:
          True
print("Enter the integer and -1 to exit")
num = int(input())
while int(num) != -1 :
     print(str(near_hundred(num)))
     num = int(input())
