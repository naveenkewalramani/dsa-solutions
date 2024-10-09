import random
class SnakeAndLadder():
    
    def __init__(self):
        self.board = [[0 for i in range(0, 100)] for j in range(0,100)]
        self.player1Location = [0]
        self.player2Location = [0]
        self.snakeLocation = [[62,5], [33,6], [49,9], [88,16], [41,20], [56,53], [98,64], [93,73], [95,75]]
        self.ladderLocation = [[2,31], [27,46], [10,32], [51,68], [61,79], [65,84], [71,91], [81,100]]
    
    def rollDice(self):
        return random.choice([1, 2, 3, 4, 5, 6])
    
    def movePlayer(self, playerNumber):
        if playerNumber == "1":
            currentLocation =  self.player1Location[-1]
        else:
            currentLocation =  self.player2Location[-1]
        stepsToMove = self.rollDice()
        newLocation = currentLocation + stepsToMove
        # check for snake:  (Assumption only one of the 2 loop will run )
        for element in self.snakeLocation:
            if element[0] == newLocation:
                newLocation = element[1]
                break
        # check for ladder: (Assumption only one of the 2 loop will run )
        for element in self.ladderLocation:
            if element[0] == newLocation:
                newLocation = element[1]
                break
        if playerNumber == "1":
            self.player1Location.append(newLocation)
        else:
            self.player2Location.append(newLocation)
        
        if newLocation == 100:
            if playerNumber == "1":
                print("Winning Player: ", playerNumber, self.player1Location)
            else:
                print("Winning Player: ", playerNumber, self.player2Location)
            exit()
        print("Game Continue")
        return
        

obj =  SnakeAndLadder()
i = 0
while(i<100):
    obj.movePlayer("1")
    obj.movePlayer("2")
    i+=1

        

