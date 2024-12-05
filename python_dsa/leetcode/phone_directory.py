# Problem Statement = https://leetcode.com/problems/design-phone-directory/

import collections

class PhoneDirectory:

    def __init__(self, maxNumbers: int):
        self.phoneDirectory = collections.OrderedDict()
        self.queue = collections.deque()
        for i in range(0, maxNumbers):
            self.phoneDirectory[i] = -1
            self.queue.append(i)

    def get(self) -> int:
        # print(self.queueImpl, self.phoneDirectory)
        if len(self.queue) == 0:
            return -1
        item = self.queue.popleft()
        self.phoneDirectory[item] = 1
        return item
        
    def check(self, number: int) -> bool:
        if self.phoneDirectory[number] == -1:
            return True
        return False
        

    def release(self, number: int) -> None:
        if self.phoneDirectory[number] == 1:
            self.phoneDirectory[number] = -1
            # insertIndex = -1
            self.queue.appendleft(number)
            # for iter in range(0, len(self.queueImpl)):
            #     if self.queueImpl[iter] >= number:
            #         insertIndex = iter
            #         break
            # self.queueImpl.insert(insertIndex, number)
        


# Your PhoneDirectory object will be instantiated and called as such:
# obj = PhoneDirectory(maxNumbers)
# param_1 = obj.get()
# param_2 = obj.check(number)
# obj.release(number)