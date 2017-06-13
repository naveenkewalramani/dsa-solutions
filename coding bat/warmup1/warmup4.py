def diff21(num1):
    if(int(num1) == 21):
        return 0
    elif(int(num1) < 21):
        return 21 - int(num1)
    elif(int(num1) > 21):
        if(int(num1)-21  > 21):
            return 2*(int(num1) - 21)
        else:
            return int(num1)-21
num1 = int(input())
while int(num1) != -1 :
    print(diff21(num1))
    num1 = int(input())
    
