#User function Template for python3

class Solution:
    ##Complete this function
    def ternarySearch(self,arr,N,K):
        return self.search(arr, K, 0, len(arr))
    
    def search(self,arr, K, left, right):
        if right >= left:
            mid1 = left + (right - left) // 3
            mid2 = right - (right - left) // 3
            if arr[mid1] == K:
                return 1
            if arr[mid2] == K:
                return 1
            if arr[mid1] > K:
                self.search(arr, K, left, mid1-1)
            elif arr[mid2] < K:
                self.search(arr, K, mid2+1,right)
            else:
                self.search(arr, K, mid1+1, mid2-1)
        return -1
                


#{ 
 # Driver Code Starts
#Initial Template for Python 3

import math
if __name__ == '__main__': 
	t = int(input())
	for _ in range(t):
		NK = input().strip().split()
		N = int(NK[0])
		K = int(NK[1])
		arr = [ int(x) for x in input().strip().split() ]
		ob=Solution()
		print(ob.ternarySearch(arr,N,K))

# } Driver Code Ends