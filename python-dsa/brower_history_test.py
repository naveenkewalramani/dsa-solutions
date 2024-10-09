import unittest
import brower_history

class TestBrowserhistory(unittest.TestCase):
    def test_happyPath(self):
        obj = brower_history.BrowserHistory("leetcode.com")
        response = [None]
        response.append(obj.visit("google.com"))
        response.append(obj.visit("facebook.com"))
        response.append(obj.visit("youtube.com"))
        response.append(obj.back(1))
        response.append(obj.back(1))
        response.append(obj.forward(1))
        response.append(obj.visit("linkedin.com"))
        response.append(obj.forward(2))
        response.append(obj.back(2))
        response.append(obj.back(7))
        self.assertEqual(response,[None,None,None,None,"facebook.com","google.com","facebook.com",None,"linkedin.com","google.com","leetcode.com"] )


if __name__ == '__main__':
    unittest.main()