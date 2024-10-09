# Problems - https://leetcode.com/problems/design-a-todo-list/

import collections
class TaskData:
    def __int__(self):
        print("initDone")
       
    def setValues(self, id, userID, taskDescription, dueDate, tags, status):
        self.ID = id
        self.UserID = userID
        self.Task = taskDescription
        self.DueDate = dueDate
        self.Tags = tags
        self.Status = status
        return self

class TodoList:

    def __init__(self):
        self.AutoIncrement = 0
        self.UserIDToTaskIDMap = {}

    def addTask(self, userId: int, taskDescription: str, dueDate: int, tags: List[str]) -> int:
        taskObj = TaskData()
        task = taskObj.setValues(self.AutoIncrement+1, userId, taskDescription, dueDate, tags, "P")
        self.AutoIncrement +=1
        if userId not in self.UserIDToTaskIDMap:
            self.UserIDToTaskIDMap[userId] = collections.OrderedDict()
            self.UserIDToTaskIDMap[userId] = {dueDate: [task]}
        else:
            currentDictionary = self.UserIDToTaskIDMap[userId]
            if dueDate not in currentDictionary:
                currentDictionary[dueDate] = [task]
            else:
                currentDueDateTask = currentDictionary[dueDate]
                currentDueDateTask.append(task)
                currentDictionary[dueDate] = currentDueDateTask
            self.UserIDToTaskIDMap[userId] = currentDictionary
        return task.ID
        

    def getAllTasks(self, userId: int) -> List[str]:
        response = []
        if userId not in self.UserIDToTaskIDMap:
            return response
        currentDictionary = self.UserIDToTaskIDMap[userId]
        dueDates = list(currentDictionary.keys())
        dueDates.sort()
        for dueDate in dueDates:
            for task in currentDictionary[dueDate]:
                if task.Status == "P":
                    response.append(task.Task)
        return response
        

    def getTasksForTag(self, userId: int, tag: str) -> List[str]:
        response = []
        if userId not in self.UserIDToTaskIDMap:
            return response
        currentDictionary = self.UserIDToTaskIDMap[userId]
        dueDates = list(currentDictionary.keys())
        dueDates.sort()
        for dueDate in dueDates:
            for task in currentDictionary[dueDate]:
                if tag in task.Tags and task.Status == "P":
                    response.append(task.Task)
        return response
        

    def completeTask(self, userId: int, taskId: int) -> None:
        if userId not in self.UserIDToTaskIDMap:
            return
        currentDictionary = self.UserIDToTaskIDMap[userId]
        for dueDate in currentDictionary:
            for task in currentDictionary[dueDate]:
                if  task.ID ==taskId and task.Status == "P" :
                    task.Status = "C"
        


# Your TodoList object will be instantiated and called as such:
# obj = TodoList()
# param_1 = obj.addTask(userId,taskDescription,dueDate,tags)
# param_2 = obj.getAllTasks(userId)
# param_3 = obj.getTasksForTag(userId,tag)
# obj.completeTask(userId,taskId)