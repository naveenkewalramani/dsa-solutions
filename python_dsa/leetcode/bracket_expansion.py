### https://leetcode.com/discuss/interview-experience/5341224/Stripe-or-Backend-Engineer-or-Bangalore-or-Jun-2024-or-Reject/
### Stripe interview old question
def bracketExpansion(string):
    leftBracketReached = False
    rightBracketReached = False
    prefixString = ""
    suffixString = ""
    tokenString=""
    for i in range(0, len(string)):
        if string[i] == "{":
            leftBracketReached = True
            if rightBracketReached == True:
                return string
        if string[i] == "}":
            rightBracketReached = True
        if string[i] != "{" and string[i] != "}":
            if leftBracketReached == False:
                prefixString += string[i]
            elif leftBracketReached == True and rightBracketReached == False:
                tokenString+=string[i]
            elif leftBracketReached == True and rightBracketReached == True:
                suffixString +=string[i]
    tokens = list(tokenString.split(","))
    if leftBracketReached == False and rightBracketReached == False:
        return string
    if len(tokens) <=1:
        return string
    response = []
    for token in tokens:
        response.append(prefixString+token+suffixString)
    return response

print(bracketExpansion("/2022/{jan,feb,march}/report"))
print(bracketExpansion("over{crowd,eager,bold,fond}ness"))
print(bracketExpansion("read.txt{,.bak}"))
print(bracketExpansion("sun{mars}rotation"))
print(bracketExpansion("minimum{}change"))
print(bracketExpansion("hello-world"))
print(bracketExpansion("hello-{-world"))
print(bracketExpansion("hello-}-weird-{-world"))


