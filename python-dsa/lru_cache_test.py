import unittest
import lru_cache

class Test_LruCache(unittest.TestCase):
    def test_happyPath_successfulPutAndGettWhenNotFull(self):
        obj = lru_cache.LruCache(4)
        obj.put(1, 1)
        response = obj.get(1)
        self.assertEqual(1, response)

    def test_happyPath_successfulPutAndGetWhenFull(self):
        obj = lru_cache.LruCache(1)
        obj.put(1, 1)
        response = obj.get(1)
        self.assertEqual(1, response)
    
    def test_happyPath_getNewKeyAndEvictedKey(self):
        obj = lru_cache.LruCache(1)
        obj.put(1, 1)
        obj.put(2, 2)
        response = obj.get(2)
        self.assertEqual(2, response)
        response = obj.get(1)
        self.assertEqual(-1, response)
    
if __name__ == '__main__':
    unittest.main()