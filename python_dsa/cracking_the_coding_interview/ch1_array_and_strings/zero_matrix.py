# Zero Matrix: Write an algorithm such that if an element in an MxN matrix is 0, its entire row and column are set to 0.

# Approach 1 - Iterate matrix, store row and column in separate array. Iterrate again and make value 0. Time = O(N2) Space = O(N)

class Solution:
    def zeroMatrix(self, matrix):
        row = []
        column = []
        for i in range(0, len(matrix)):
            for j in range(0, len(matrix[0])):
                if matrix[i][j] == 0:
                    row.append(i)
                    column.append(j)
        for i in row:
            for j in range(0, len(matrix[0])):
                matrix[i][j] = 0
        for j in column:
            for i in range(0, len(matrix)):
                matrix[i][j] = 0
        return matrix

if __name__ == "__main__":
    obj = Solution()
    print(obj.zeroMatrix([[1,1], [2,2]]))
    print(obj.zeroMatrix([[0,1], [1,1]]))