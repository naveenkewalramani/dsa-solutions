#Given an array length 1 or more of ints, return the difference between the largest and smallest values in the array. Note: the built-in min(v1, v2) and max(v1, v2) functions return the smaller or larger of two values.
#big_diff([10, 3, 5, 6]) → 7
#big_diff([7, 2, 10, 9]) → 8
#big_diff([2, 10, 7, 2]) → 8
def big_diff(array):
     array = sorted(array)
     diff = int(array[-1]) - int(array[0])
     return diff
array = []
print('Enter the number or -1 to exit')
num  = int(input())
while num != -1:
     array.append(num)
     print('Enter the numebr or -1 to exit')
     num = int(input())
print(big_diff(array))
