class Solution:
    def longestSubstring(self, inputStr):
        l = 0
        r = 0
        result = 0
        n = len(inputStr)
        chars = [0 for _ in range(0, 257)]
        while(r < n):
            t1 = ord(inputStr[r])
            chars[t1]+=1
            while(chars[t1]>1):
                t2 = ord(inputStr[l])
                chars[t2]-=1
                l+=1
            result = max(result, r - l +1)
            r+=1
        return result
    
if __name__ == "__main__":
    obj = Solution()
    print(obj.longestSubstring("abcabcabc"))
    print(obj.longestSubstring("abcadea"))