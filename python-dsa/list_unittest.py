import unittest
import list

class TestList(unittest.TestCase):
    
    def test_add(self):
        result = list.add(20, 39)
        self.assertEqual(result, 59)

    def test_subtract(self):
         result = list.subtract(20, 39)
         self.assertEqual(result, -19) 

if __name__ == '__main__':
    unittest.main()