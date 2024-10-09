import unittest
class TestClass(unittest.TestCase):
    def test_happyPath(self):
        self.assertEqual(10, 10)

if __name__ == "__main__":
    unittest.main()