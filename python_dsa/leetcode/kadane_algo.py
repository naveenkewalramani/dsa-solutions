# largest sum contigous subarry
def kadaneAlogrithm(arr):
    maxEndingHere = 0
    maxEndingSoFar = 0
    start = 0
    end = 0
    s = 0
    for iter in range(0, len(arr)):
        maxEndingHere = maxEndingHere + arr[iter]
        if maxEndingHere < 0:
            s = iter+1
            maxEndingHere = 0
        
        if maxEndingHere > maxEndingSoFar:
            maxEndingSoFar = maxEndingHere
            start = s
            end = iter
    return maxEndingSoFar