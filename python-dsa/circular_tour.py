'''
    lis[][0]:Petrol
    lis[][1]:Distance
'''

class Solution:
    
    #Function to find starting point where the truck can start to get through
    #the complete circle without exhausting its petrol in between.
    def tour(self,lis, n):
        s = 0
        for i in range(0,n):
            s = s + lis[i][0] - lis[i][1]
        
        if s < 0:
            return -1
        
        index = 0
        remaining = 0
        for i in range (0, n):
            remaining = remaining + lis[i][0] - lis[i][1]
            if remaining < 0:
                remaining = 0
                index = i +1
        return index
#{ 
 # Driver Code Starts
if __name__ == '__main__': 
    t = int(input())
    for i in range(t):
        n = int(input())
        arr=list(map(int, input().strip().split()))
        lis=[]
        for i in range(1, 2*n, 2):
            lis.append([ arr[i-1], arr[i] ])
        print(Solution().tour(lis, n))
    # Contributed by: Harshit Sidhwa
# } Driver Code Ends