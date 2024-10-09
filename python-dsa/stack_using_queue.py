class MyStack:

    def __init__(self):
        self.mainQueue = []
        self.tempQueue = []
        

    def push(self, x: int) -> None:
        self.mainQueue.append(x)
        

    def pop(self) -> int:
        while(len(self.mainQueue)!=1):
            self.tempQueue.append(self.mainQueue.pop(0))
        item = self.mainQueue.pop(0)
        self.mainQueue = self.tempQueue
        self.tempQueue = []
        return item
        

    def top(self) -> int:
        return self.mainQueue[-1]
        

    def empty(self) -> bool:
        if len(self.mainQueue) == 0:
            return True
        return False
        


# Your MyStack object will be instantiated and called as such:
# obj = MyStack()
# obj.push(x)
# param_2 = obj.pop()
# param_3 = obj.top()
# param_4 = obj.empty()