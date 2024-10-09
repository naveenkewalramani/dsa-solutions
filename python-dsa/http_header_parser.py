class HTTPHeaderParser():
    def __init__(self, languagesSupported):
        self.languagesSupported = {}
        for item in languagesSupported:
            self.languagesSupported[item] = True

    def checkAndReturn(self, headers):
        response = []
        for header in headers:
            try:
                if self.languagesSupported[header]:
                    response.append(header)
            except:
                print("skipping, not supported", header)
        return response
    
# obj = HTTPHeaderParser(["en-US","fr-CA", "mr-IN", "hi-IN"])
# print(obj.checkAndReturn(["fr-CA", "en-US","mr-IN", "123"]))
# print(obj.checkAndReturn(["fr-CA", "fr-ee","mr-IN", "123"]))