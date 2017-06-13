#Given an array of ints, return a new array length 2 containing the first and last elements from the original array. The original array will be length 1 or more.
#make_ends([1, 2, 3]) → [1, 3]
#make_ends([1, 2, 3, 4]) → [1, 4]
#make_ends([7, 4, 6, 2]) → [7, 2]

def make_ends(a):
     b = []
     b.append(int(a[0]))
     b.append(int(a[-1]))
     return b
a = []
print('Enter the number or -1 to exit')
num = int(input())
while int(num) != -1:
     a.append(int(num))
     print('Enter the number or -1 to exit')
     num = int(input())
print(make_ends(a))
