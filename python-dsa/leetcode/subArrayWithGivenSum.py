class SubArrayWithGivenSum():
    def result(self, arr, key):
        hashmap = {}
        currSum = 0
        for i in range(0, len(arr)):
            currSum += arr[i]
            if currSum == key:
                return [0, i]
            if (currSum - key) in hashmap:
                return [hashmap[currSum -key]+1, i]
            hashmap[currSum] = i
        return [-1, -1] 