class HashMap():
    def __init__(self):
        self.ValueList = [-1 for i in range(0, 10**3)]

    def put(self, key, value):
       self.ValueList[key] = value

    def get(self, key):
        value  = self.ValueList[key]
        if value == -1:
            return -1
        return value    

    def remove(self, key):
        self.ValueList[key] = -1

obj = HashMap()
# obj.put(1, 1)
# obj.put(2, 2)
# print(obj.get(1))
# print(obj.get(3))
# obj.put(2, 1)
# print(obj.get(2))
# obj.remove(2)
# print(obj.get(2))

obj.put(1,1)
obj.put(2,2)
print(obj.get(2))

