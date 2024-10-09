class WordDictionary:

    def __init__(self):
        self.WordList = {}    

    def addWord(self, word: str) -> None:
        self.WordList[word] = True

    def search(self, word: str) -> bool:
        if len(self.WordList) == 0:
            # no data in the dictionary
            return False
        try:
            value = self.WordList[word]
            # if word matches directly
            return True
        except:
            for key in self.WordList.keys():
                if len(key) != len(word):
                    # case where the size does not match
                    continue
                flag = True
                for iter in range(0, len(key)):
                    if word[iter] != "." and word[iter] != key[iter]:
                        flag = False
                return flag
            return False
        


# Your WordDictionary object will be instantiated and called as such:
# obj = WordDictionary()
# obj.addWord(word)
# param_2 = obj.search(word)