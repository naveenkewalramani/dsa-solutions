import unittest
import bubble_sort

class TestBubbleSort(unittest.TestCase):
    def test_happyPath(self):
        input = [3, 4, 5, 2, 1, 3]
        expectedResponse = [1, 2, 3, 3, 4, 5]
        self.assertEqual(expectedResponse, bubble_sort.bubbleSort(input))

if __name__ == '__main__':
    unittest.main()