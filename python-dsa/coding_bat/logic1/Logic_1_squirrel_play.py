#The squirrels in Palo Alto spend most of the day playing. In particular, they play if the temperature is between 60 and 90 (inclusive). Unless it is summer, then the upper limit is 100 instead of 90. Given an int temperature and a boolean is_summer, return True if the squirrels play and False otherwise.
#squirrel_play(70, False) → True
#squirrel_play(95, False) → False
#squirrel_play(95, True) → True
def squirrel_play(num,logic):
     if logic  == False and num>=60 and num<=90:
          return True
     elif logic  == True and num >= 60 and num<=100:
          return True
     else:
          return False
print('Enter the number')
num = int(input())
print('Enter the logic is_Summer')
logic = bool(input())
print(squirrel_play(num,logic))
