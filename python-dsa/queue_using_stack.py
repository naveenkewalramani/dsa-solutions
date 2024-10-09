class MyQueue:

    def __init__(self):
        self.mainStack = []
        self.tempStack = []
        self.peekElement = 0
        self.size  = 0

    def push(self, x: int) -> None:
        if len(self.mainStack) == 0:
            self.mainStack.append(x)
            self.peekElement = x
            self.size +=1
        else:
            self.mainStack.append(x)
            self.size +=1

    def pop(self) -> int:
        while (len(self.mainStack) != 1):
            self.tempStack.append(self.mainStack.pop())
        item = self.mainStack.pop()
        self.size -=1
        if len(self.tempStack) != 0:
            self.peekElement = self.tempStack[-1]
            while (len(self.tempStack)!=0):
                self.mainStack.append(self.tempStack.pop())
        else:
            self.peekElement = 0
        return item

    def peek(self) -> int:
        return self.peekElement
        

    def empty(self) -> bool:
        if self.size == 0:
            return True
        return False
        


# Your MyQueue object will be instantiated and called as such:
# obj = MyQueue()
# obj.push(x)
# param_2 = obj.pop()
# param_3 = obj.peek()
# param_4 = obj.empty()