import unittest
import kadane_algo

class TestKadaneAlgo(unittest.TestCase):
    def test_happyPath_allPositive(self):
        self.assertEqual(kadane_algo.kadaneAlogrithm([1,2,3,4]), 10)
    
    def test_happyPath_positiveAndNegative(self):
        self.assertEqual(kadane_algo.kadaneAlogrithm([1,2,-2,4]), 5)

if __name__ == '__main__':
    unittest.main()