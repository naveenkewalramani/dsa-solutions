import unittest
import gcd

class TestGcd(unittest.TestCase):
    def test_gcdhappypath(self):
        result = gcd.gcd(10, 20)
        self.assertEqual(10, result)

if __name__ == "__main__":
    unittest.main()