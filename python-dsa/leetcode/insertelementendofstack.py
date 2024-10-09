
class Stack:
    def insertElementToStack(self, st:list, x)->list:
        newStack = []
        while(len(st) != 0):
            newStack.append(st.pop())
        st.append(x)
        while(len(newStack) != 0):
            st.append(newStack.pop())
        return st