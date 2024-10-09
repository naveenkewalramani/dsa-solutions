class User():
    def __init__(self,id):
        self.id = id
        self.balance = 0

class Expense():
    def __init__(self, amount, paidBy, participants):
        self.amount = amount
        self.paidBy = paidBy
        self.participants = participants
    
    def splitExpense(self):
        splitAmount = self.amount / len(self.participants)
        for participant in self.participants:
            if participant.id != self.paidBy.id:
                participant.balance -= splitAmount
                self.paidBy.balance += splitAmount
    
class Splitwise():
    def __init__(self):
        self.users = []
    
    def addUser(self, id):
        obj = User(id)
        self.users.append(obj)
        return obj
    
    def add_expense(self, amount, paid_by, pariticipants):
        expense = Expense(amount, paid_by, pariticipants)
        expense.splitExpense()
    
    def showBalance(self):
        for user in self.users:
            print("User: ", user.id, ", balance: ", user.balance)

if __name__ == "__main__":
    splitwise = Splitwise()
    user1 = splitwise.addUser("Naveen")
    user2 = splitwise.addUser("Ayushi")
    user3 = splitwise.addUser("Interviewer")
    splitwise.showBalance()
    splitwise.add_expense(200, user1, [user2])
    splitwise.add_expense(100, user1, [user2, user3])
    splitwise.add_expense(50, user2, [user1, user3])
    splitwise.add_expense(20, user3, [user1])
    splitwise.showBalance()

