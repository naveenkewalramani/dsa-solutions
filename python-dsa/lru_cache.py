class LruCache():
    def __init__(self, size):
        self.map = {}
        self.lis =[]
        self.size = size
    
    def put(self, key, value):
        if key not in self.map:
            # element not present
            if len(self.lis) == self.size:
                # need to delete last
                lastElement = self.lis[-1]
                self.lis.pop()
                del self.map[lastElement]
        else:
            self.lis.remove(key)
        self.map[key] = value
        self.lis.insert(0,key)
    
    def get(self, key):
        if key in self.map:
            return self.map[key]
        return -1
    

# obj = LruCache(4)
# obj.put("key1", "value1")
# obj.put("key2", "value2")
# print(obj.get("key1"))
# obj.put("key3", "value3")
# obj.put("key4", "value4")
# print(obj.get("key3"))
# obj.put("key3", "value3.1")
# obj.put("key5", "value5")
# print(obj.get("key3"))
# print(obj.get("key5"))
# print(obj.get("key1"))