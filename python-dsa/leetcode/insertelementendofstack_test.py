import unittest
import insertelementendofstack


class TestStack(unittest.TestCase):
    def test_insertElementToStack_happyPath(self):
        newObj = insertelementendofstack.Stack()
        element = 20
        existingStack = [3, 4, 5,132]
        expectedStack = [20, 3, 4, 5, 132]
        self.assertEqual(newObj.insertElementToStack(existingStack, element), expectedStack)
    
    def test_insertElementToStack_emptyStack(self):
        newObj = insertelementendofstack.Stack()
        element = 20
        existingStack = []
        expectedStack = [20]
        self.assertEqual(newObj.insertElementToStack(existingStack, element), expectedStack)

if __name__ == '__main__':
    unittest.main()