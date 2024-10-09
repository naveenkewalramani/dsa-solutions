class MinStack:

    def __init__(self):
        self.stackList = []
        
    def push(self, val: int) -> None:
        if len(self.stackList) == 0:
            self.stackList.append([val, val])
        else:
            minElement = min(self.stackList[-1][1], val)
            self.stackList.append([val, minElement])

    def pop(self) -> None:
        self.stackList.pop()

    def top(self) -> int:
        return self.stackList[-1][0]

    def getMin(self) -> int:
        return self.stackList[-1][1]
        


# Your MinStack object will be instantiated and called as such:
# obj = MinStack()
# obj.push(val)
# obj.pop()
# param_3 = obj.top()
# param_4 = obj.getMin()