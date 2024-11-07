# String Compression: Implement a method to perform basic string compression using the counts
# of repeated characters. For example, the string aabcccccaaa would become a2blc5a3. If the
# "compressed" string would not become smaller than the original string, your method should return
# the original string. You can assume the string has only uppercase and lowercase letters (a - z).


class Solution:
    def stringCompressor(self, inputStr):
        prevChar = inputStr[0]
        count = 1
        outputStr = ""
        n = len(inputStr)
        for i in range(1, n):
            if inputStr[i] == prevChar:
                count+=1
            else:
                outputStr += prevChar
                outputStr += str(count)
                prevChar = inputStr[i]
                count = 1
        outputStr += prevChar
        outputStr += str(count)
        if len(outputStr) > len(inputStr):
            return inputStr
        return outputStr

if __name__ == "__main__":
    obj = Solution()
    print("Test Case 1: ",  obj.stringCompressor("aabbccAABBCCaaBBccAAAAaa") )
    print("Test Case 2: ",  obj.stringCompressor("aabbcc") )
    print("Test Case 3: ",  obj.stringCompressor("ccccccB") )

