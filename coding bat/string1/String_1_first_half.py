#Given a string of even length, return the first half. So the string "WooHoo" yields "Woo".

#first_half('WooHoo') → 'Woo'
#first_half('HelloThere') → 'Hello'
#first_half('abcdef') → 'abc'

def first_half(string):
     num = len(string)
     return string[0:(num//2)]
print('Enter the string')
string = str(input())
print(str(first_half(string)))
