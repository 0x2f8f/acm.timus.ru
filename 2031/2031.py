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

variants = ["66", "96", "16", "06", "68", "88", "98", "18", "08", "69", "89", "99", "19", "09", "61", "81", "91", "11", "01"]
count = int(s[0])

if count == 1:
    print("01")
    sys.exit()
    
if count == 2:
    print("11 01")
    sys.exit()

if count == 3:
    print("16 06 68")
    sys.exit()

if count == 4:
    print("16 06 68 88")
    sys.exit()
    
print("Glupenky Pierre");
