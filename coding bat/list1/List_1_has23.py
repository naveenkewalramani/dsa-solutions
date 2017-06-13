#Given an int array length 2, return True if it contains a 2 or a 3.
#has23([2, 5]) → True
#has23([4, 3]) → True
def has23(a):
     if int(a[0]) == 2 or int(a[1]) == 2 or int(a[1])==3 or int(a[0])==3 :
          return True
     else:
          return False
print(str(has23([2,5])))
print(str(has23([4,3])))
print(str(has23([4,5])))
