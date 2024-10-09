class Solution:
    def calcEquation(self, equations: List[List[str]], values: List[float], queries: List[List[str]]) -> List[float]:

        graph = defaultdict(defaultdict)

        def bfs(curr_node, target_node, acc_product, visited):
            ret = -1.0
            queue=[]
            queue.append([curr_node, acc_product])
            visited.add(curr_node)
            while(len(queue) != 0):
                item = queue.pop() 
                if item[0] == target_node:
                    return item[1]
                neighbors = graph[item[0]]
                for neighbor, value in neighbors.items():
                    if neighbor in visited:
                        continue
                    queue.append([neighbor, item[1] * value])
                    visited.add(neighbor)
            return ret

        # Step 1). build the graph from the equations
        for (dividend, divisor), value in zip(equations, values):
            # add nodes and two edges into the graph
            graph[dividend][divisor] = value
            graph[divisor][dividend] = 1 / value

        # Step 2). Evaluate each query via backtracking (DFS)
        #  by verifying if there exists a path from dividend to divisor
        results = []
        for dividend, divisor in queries:
            if dividend not in graph or divisor not in graph:
                # case 1): either node does not exist
                ret = -1.0
            elif dividend == divisor:
                # case 2): origin and destination are the same node
                ret = 1.0
            else:
                visited = set()
                ret = bfs(dividend, divisor, 1, visited)
            results.append(ret)

        return results