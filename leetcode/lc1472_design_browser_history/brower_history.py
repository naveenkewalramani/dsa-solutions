class BrowserHistory:

    def __init__(self, homepage: str):
        self.history = []
        self.current = homepage
        self.future = []

    def visit(self, url: str) -> None:
        self.history.append(self.current)
        self.current = url
        self.future = []
        

    def back(self, steps: int) -> str:
        while steps >0 and len(self.history) != 0:
            self.future.append(self.current)
            self.current = self.history.pop()
            steps -= 1
        return self.current
        

    def forward(self, steps: int) -> str:
        while steps > 0 and len(self.future):
            self.history.append(self.current)
            self.current = self.future.pop()
            steps -=1
        return self.current
