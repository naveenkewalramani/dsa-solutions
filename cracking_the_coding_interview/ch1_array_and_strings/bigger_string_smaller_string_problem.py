def main(b, s):
    position = []
    x = list(b)
    y = list(s)
    for i in range(0, len(b)-len(s)): # O(B - S)
        if compareIfPermutation(x[i:i+len(s)], y):
            position.append(i)
    return position

def compareIfPermutation(m, n): # O(SlogS)
    m.sort()
    n.sort()
    return m == n

if __name__ == "__main__":
    print(main("abahabababa", "ab"))
