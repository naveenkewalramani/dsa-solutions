#Given an array of ints length 3, return the sum of all the elements.
#sum3([1, 2, 3]) → 6
#sum3([5, 11, 2]) → 18
#sum3([7, 0, 0]) → 7
def sum3(a):
     return int(a[0]) + int(a[1]) + int(a[2])
count=0
a=[]
while count != 3 :
     print('Enter the number')
     num = int(input())
     a.append(num)
     count = count+1
print(str(sum3(a)))
