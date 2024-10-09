import unittest
import computeprice

class Test_computePrice(unittest.TestCase):
    def test_computePrice(self):
        self.assertEqual(computeprice.computePrice(20, 3, 6), 17)

if __name__ == '__main__':
    unittest.main()