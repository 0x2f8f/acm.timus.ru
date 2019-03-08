import pprint
import sys
import math

"""
fin = open("a.in")
s = fin.read().split('\n')
#s = fin.read()
#s = fin.readline().strip()
"""
s = []
s.append(input())

#pprint.pprint(s)
#sys.exit()
count = int(s[0])
n3 = 0
n4 = 0
n5 = 0
summ = 0
i = 1
while i <= count:
    s.append(input())
    grade = int(s[i])
    summ += grade
    if grade == 3:
        n3 += 1
    elif grade == 4:
        n4 += 1
    elif grade == 5:
        n5 += 1
    i += 1

if n5 == count:
    print ("Named")
else:
    if n3 > 0:
        print("None")
    else:
        if summ/count >= 4.5:
            print("High")
        else:
            print("Common")
