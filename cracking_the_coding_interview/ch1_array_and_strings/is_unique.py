def isUnique(s:str):
    l = list(s)
    l.sort()
    for  i in range(0, len(l)-1):
        if l[i] == l[i+1]:
            return False
    return True

if __name__ == "__main__":
    print(isUnique("abcdefghijk"))
    print(isUnique("baa"))