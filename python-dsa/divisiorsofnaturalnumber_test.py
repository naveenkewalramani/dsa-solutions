import unittest
import divisiorsofnaturalnumber

class TestDivisorOfNaturalNumber(unittest.TestCase):
    def test_happyPath_naturalNumber(self):
        result = divisiorsofnaturalnumber.divisiorsOfNaturalNumber(10)
        result.sort()
        self.assertEqual([1, 2, 5, 10], result)
    
    def test_happyPath_numberzero(self):
        result = divisiorsofnaturalnumber.divisiorsOfNaturalNumber(0)
        result.sort()
        self.assertEqual([], result)
    
    def test_happyPath_numberone(self):
        result = divisiorsofnaturalnumber.divisiorsOfNaturalNumber(1)
        result.sort()
        self.assertEqual([1], result)


if __name__ == '__main__':
    unittest.main()