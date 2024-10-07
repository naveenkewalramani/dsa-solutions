#Given an "out" string length 4, such as "<<>>", and a word,
#return a new string where the word is in the middle of the out string,
#e.g. "<<word>>".

#make_out_word('<<>>', 'Yay') → '<<Yay>>'
#make_out_word('<<>>', 'WooHoo') → '<<WooHoo>>'
#make_out_word('[[]]', 'word') → '[[word]]'
def make_out_word(string1,string2):
     return string1[0:2] + string2 + string1[2:4]
print('Enter the two string')
string1 = str(input())
string2 = str(input())
print(str(make_out_word(string1,string2)))
