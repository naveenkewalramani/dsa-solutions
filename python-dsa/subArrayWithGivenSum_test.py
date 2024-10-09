import unittest
import subArrayWithGivenSum
class TestSubArrayWithGivenSum(unittest.TestCase):
    def test_result_elmentEqualSum(self):
        obj = subArrayWithGivenSum.SubArrayWithGivenSum()
        response = obj.result([1,7,3,4,5], 5)
        self.assertEqual([4,4], response)

    def test_result_notPossible(self):
        obj = subArrayWithGivenSum.SubArrayWithGivenSum()
        response = obj.result([1,7,3,4,5],95)
        self.assertEqual([-1,-1], response)

if __name__ == '__main__':
    unittest.main()