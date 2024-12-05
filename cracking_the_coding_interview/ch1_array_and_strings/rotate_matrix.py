class Solution:
    def rotateMatrix(self, grid):
        n = len(grid)
        m = len(grid[0])
        for i in range(n//2+n%2):
            for j in range(m//2):
                temp = grid[n-1-j][i]
                grid[n-1-j][i] = grid[n-1-i][n-1-j]
                grid[n-1-i][n-1-j] = grid[j][n-1-i]
                grid[j][n-1-j] = grid[i][j]
                grid[i][j] = temp
        return grid

if __name__ == "__main__":
    obj = Solution()
    print(obj.rotateMatrix([[0,0, 0],[1,1,1]]))