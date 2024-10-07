from collections import defaultdict

def palindromePermutation(inputStr):
    # considerations - space to be ignored.
    # case insensitive 
    # considering all the characters even special
    hashmap = defaultdict()
    n = len(inputStr)
    for i in inputStr:
        if i.isalpha():
            value = hashmap.get(i, 0)
            value+=1
            hashmap[i] = value
    diff = 0
    for key, value in hashmap.items():
        if value %2 != 0:
            diff +=1
        if diff>1:
            return False
    return True

if __name__ == "__main__":
    print(palindromePermutation("au uda"))