import unittest
import selectionsort

class TestSelectionSort(unittest.TestCase):
    def test_selectionsort(self):
        data = [19, 0, 29, 19]
        expectedResponse = [0, 19, 19, 29]
        self.assertEqual(expectedResponse, selectionsort.selectionSort(data))
    
    def test_selectionsort2(self):
        data = [19, 0, 0, 19, -1]
        expectedResponse = [-1, 0, 0, 19, 19]
        self.assertEqual(expectedResponse, selectionsort.selectionSort(data))
    
if __name__ == '__main__':
    unittest.main()