def oneWay(s1, s2):
    if s1 == s2:
        return True
    if len(s1) == len(s2):
        diff = 0
        for i in range(0, len(s1)):
            if s1[i]!= s2[i]:
                diff+=1
        if diff > 1:
            return False
        return True
    else:
        p1, p2 = 0, 0 
        while(p1<len(s1) and p2<len(s2)):
            if s1[p1] != s2[p2]:
                p1+=1
            elif s1[p1] == s2[p2]:
                p1+=1
                p2+=1
        if p2 == len(s2):
            return True
        return False

print(oneWay("pale", "ple"))
print(oneWay("pales", "pale"))
print(oneWay("pale", "bale"))
print(oneWay("pale", "bake"))