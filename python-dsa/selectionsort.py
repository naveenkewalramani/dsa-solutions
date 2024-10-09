def selectionSort(arr):
    for iter in range (0, len(arr)-1):
        minIndex = iter
        for iter2 in range(iter+1, len(arr)):
            if arr[iter2] < arr[minIndex]:
                minIndex = iter2
        if iter != minIndex:
            arr[iter], arr[minIndex] = arr[minIndex], arr[iter]
    return arr