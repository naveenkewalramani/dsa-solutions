import unittest
import oddoreven

class Test_oddoreven(unittest.TestCase):
    def test_oddoreven_1(self):
        resp = oddoreven.OddOrEven(20)
        self.assertEqual(resp, True)
    
    def test_oddoreven_2(self):
        resp = oddoreven.OddOrEven(31)
        self.assertEqual(resp, False)

if __name__ == '__main__':
    unittest.main()