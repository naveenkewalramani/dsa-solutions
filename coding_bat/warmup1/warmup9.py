def missing_char(string,num):
     print(string[0:int(num)] + string[int(num)+1:len(string)])
missing_char('kitten',1)
missing_char('kitten',0)
missing_char('kitten',4)
