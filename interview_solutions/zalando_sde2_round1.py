# Write a function Solution that, given a string S of length N,
#      returns any palindrome which can be obtained by replacing all of the question marks in S by lowercase letters ('a' − 'z'). 
#      If no palindrome can be obtained, the function should return the string "NO".

# A palindrome is a string that reads the same both forwards and backwards. Some examples of palindromes are: "kayak", "radar", "mom".

# Examples:

# 1. Given S = "?ab??a", the function should return "aabbaa".

# 2. Given S = "bab??a", the funct

# Examples:

# 1. Given S = "?ab??a", the function should return "aabbaa".

# 2. Given S = "bab??a", the function should return "NO".

# 3. Given S = "?a?", the function may return "aaa". It may also return "zaz", among other possible answers.
#      The function is supposed to return only one of the possible answers.

# Assume that:

# N is an integer within the range [1..1,000];
# string S consists only of lowercases letters ('a' − 'z') or '?'.

def replaceQuestionMark(inputStr):
    inputStrToList = list(inputStr)
    n = len(inputStr)
    for i in range(n//2):
        left = inputStrToList[i]
        right = inputStrToList[n-i-1]
        if left == "?" and right == "?":
            inputStrToList[i] = inputStrToList[n-i-1] = "a"
        elif left == "?":
            inputStrToList[i] = right
        elif right == "?":
            inputStrToList[n-i-1] = left
        elif left != right:
            return "NO"

    if n %2 == 1 and inputStrToList[n//2] == "?":
        inputStrToList[n//2] = "a"
    
    return ''.join(inputStrToList)

def formatString(inputStr):
    s = list(inputStr)
    response = []
    for element in s:
        if element == "?":
            response.append(element)
        elif element.isalpha():
            response.append(element.lower())
    return "".join(response)

def buildPalindrome(inputStr):
    parsedStr = formatString(inputStr)
    updatedStr = replaceQuestionMark(parsedStr)
    if updatedStr == "NO":
        return len(inputStr) // 2 * "no"
    else:
        responseStr = ""
        index = 0
        for i in inputStr:
            if i == "?" or i.isalpha():
                responseStr += updatedStr[index]
                index+=1
            else:
                responseStr += i
    return responseStr
            

print(replaceQuestionMark("?ab??a"))
print(replaceQuestionMark("bab??a"))
print(replaceQuestionMark("?a?"))

print("Output from second set of test cases")
print(buildPalindrome("a?-b=BA"))
print(buildPalindrome("?+B?b+--a"))
print(buildPalindrome("?+cd+a"))
print(buildPalindrome("1a?-b=BA"))
print(buildPalindrome("1?"))

















# This iteration's string consists of lowercase and uppercase letters special characters.
# The algorithm must work case insensitive, and special characters are useless but the first logic is still valid so ? 
# should be ignored as a special character. 
# The result must be lowercase. If the string is not a palindrome, return a string with the same length and fill it with "no."

# Example:
# - "a?-b=BA"    => "ab-b=ba"
# - "?+B?b+--a" => "a+bab+--a"
# - "?+cd+a"      => "nonono"
# - "1a?-b=BA"    => "1ab-b=ba"  OR no