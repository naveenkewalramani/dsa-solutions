# ProblemLink = https://leetcode.com/problems/zigzag-iterator/

class ZigzagIterator:
    def __init__(self, v1: List[int], v2: List[int]):
        self.flattenVector = []
        maxLength = max(len(v1), len(v2))
        for iter in range(0, maxLength):
            if iter < len(v1):
                self.flattenVector.append(v1[iter])
            if iter < len(v2):
                self.flattenVector.append(v2[iter])
        print(self.flattenVector)

    def next(self) -> int:
        item = self.flattenVector.pop(0)
        return item
        

    def hasNext(self) -> bool:
        if len(self.flattenVector) == 0:
            return False
        return True
        

# Your ZigzagIterator object will be instantiated and called as such:
# i, v = ZigzagIterator(v1, v2), []
# while i.hasNext(): v.append(i.next())