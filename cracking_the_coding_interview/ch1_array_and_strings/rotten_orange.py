from collections import deque

def rottenOrange(grid):
    queue = deque()
    response = []
    if not grid:
        return response
    n = len(grid)
    m = len(grid[0])
    noOfOne =  0
    for i in range(0,n):
        for j in range(0,m):
            if grid[i][j] == 2:
                queue.append([i,j])
            elif grid[i][j] == 1:
                noOfOne+=1
    if len(queue) == 0 and noOfOne !=0:
        return -1
    if len(queue) == 0:
        return 0
    time = -1
    moves = [[0,1], [0,-1], [1,0], [-1,0]]

    while(len(queue) !=0):
        size = len(queue)
        for i in range(0, size):
            location = queue.popleft()
            for move in range(moves):
                r, c = location[0] + move[0], location[1] + move[1]
                if r < 0 or r >= n or c < 0 or c >=m:
                    continue
                if grid[r][c] !=1:
                    continue
                grid[r][c] = 2
        time +=1
    
    for i in range(n):
        for j in range(m):
            if grid[i][j] == 1:
                return -1
    return time