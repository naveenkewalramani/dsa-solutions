def checkPermutation(s1, s2):
    # Best case both strings are equal
    if s1 == s2:
        return True
    m = len(s1)
    n = len(s2)
    # string have unequal length
    if m != n:
        return False
    # convert to list, sort and compare. Space = O(max(n,m)), Time = O(N*logN)
    # use hashmap, Space = O(1), Time = O(max(m,n))
    xor = ord(s1[0])
    for i in range(1, m):
        xor = xor ^ ord(s1[i])
    for i in range(0, n):
        xor = xor ^ ord(s2[i])
    return xor == 0

if __name__ == "__main__":
    print(checkPermutation("abc", "cba"))
    print(checkPermutation("abc", "cbaa"))
    print(checkPermutation("abc", "abc"))
    print(checkPermutation("abc", "bbc"))