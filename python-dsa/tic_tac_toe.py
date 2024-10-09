class TicTacToe():
    def __init__(self):
        self.board = [[' ' for index in range(0, 3)] for index in range(0, 3)]
        self.lastPlayer = ""
    
    def validMove(self, row, column)->bool:
        if row >=0 and row < len(self.board) and column >=0 and column < len(self.board[1]) and self.board[row][column] == ' ':
            return True
        return False
    
    def checkWinner(self):
        for i in range(0, 3):
            if self.board[i][0] == self.board[i][1] == self.board[i][2] != ' ':
                return True
            if self.board[0][i] == self.board[1][i] == self.board[2][i] != ' ':
                return True
        if self.board[0][0] ==  self.board[1][1] ==  self.board[2][2] != ' ':
            return True
        if self.board[0][2] ==  self.board[1][1] ==  self.board[2][0] != ' ':
            return True
        return False
    
    def makeMove(self, player, row, column):
        if player == "A":
            self.board[row][column] = "X"
        else:
            self.board[row][column] = "O"
        self.lastPlayer =  player
    
    def isBoardFull(self):
        flag = True
        for i in range (0, 3):
            for j in range(0, 3):
                if self.board[i][j] == ' ':
                    flag = False
                    return flag
        return flag
    
    def startTurn(self, player, row, column):
        if self.lastPlayer == player:
            print("error cannot repeat turn")
            return 
        if self.isBoardFull():
            print("board full cannot move further")
            return
        if self.validMove(row, column) != True:
            print("error unable to move, not a valid move")
            return
        self.makeMove(player, row, column)
        print("Valid Turn", player, row, column)
        if self.checkWinner():
            print(self.lastPlayer + " is the winner")


obj = TicTacToe()
print(obj.board)
obj.startTurn("A", -1,-1)
obj.startTurn("A", 3,3)
obj.startTurn("A",0 ,0)
obj.startTurn("A",0 ,1)
obj.startTurn("B",0 ,0)
obj.startTurn("B",0 ,1)
print(obj.board)

