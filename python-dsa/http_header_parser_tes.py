import unittest
import http_header_parser

class Test_HttpParser(unittest.TestCase):
    def test_happyPath(self):
        obj = http_header_parser.HTTPHeaderParser(["en-US","fr-CA", "mr-IN", "hi-IN"])
        expectedResponse = ['fr-CA', 'en-US', 'mr-IN']
        userInput = ["fr-CA", "en-US","mr-IN", "123"]
        self.assertEqual(expectedResponse, obj.checkAndReturn(userInput))
    
    def test_happyPath_ordercheck(self):
        obj = http_header_parser.HTTPHeaderParser(["en-US","fr-CA", "mr-IN", "hi-IN", "fr-ee"])
        expectedResponse = ['fr-CA', 'en-US',"fr-ee" ,'mr-IN']
        userInput = ["fr-CA", "en-US", "fr-ee","mr-IN", "123"]
        self.assertEqual(expectedResponse, obj.checkAndReturn(userInput))
    
    def test_happyPath_supportNoLanguage(self):
        obj = http_header_parser.HTTPHeaderParser([])
        expectedResponse = []
        userInput = ["fr-CA", "en-US", "fr-ee","mr-IN", "123"]
        self.assertEqual(expectedResponse, obj.checkAndReturn(userInput))
    
    def test_happyPath_emptyUserInput(self):
        obj = http_header_parser.HTTPHeaderParser(["en-US","fr-CA", "mr-IN", "hi-IN", "fr-ee"])
        expectedResponse = []
        userInput = []
        self.assertEqual(expectedResponse, obj.checkAndReturn(userInput))
    

if __name__ =="__main__":
    unittest.main()