#Given two int values, return their sum. Unless the two values are the same, then return double their sum.
#sum_double(1, 2) → 3
#sum_double(3, 2) → 5
#sum_double(2, 2) → 8

def sum_double(num1,num2):
    if(int(num1) == int(num2)):
        return int(num1)*2 + int(num2)*2
    else:
        return int(num1) + int(num2)

while True:
    num1 = int(input())
    num2 = int(input())
    print(sum_double(num1,num2))
