def urlify(inputStr, n):
    # Sol1 - trim the string, convert to list and replace all the string with char
    # Sol2 - trim the string, iterate and replace values in another strin
    inputStr  = inputStr.strip()
    outputStr = ""
    i = 0
    while(i < n):
        if inputStr[i] != " ":
            outputStr += inputStr[i]
        else:
            outputStr += "%20"
        i+=1
    return outputStr

if __name__ =="__main__":
    print(urlify("Mr John  Smith    ", 14))
    print(urlify("Mr John Smith    ", 13))