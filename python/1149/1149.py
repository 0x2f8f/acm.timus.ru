import pprint
import sys
import math


def funSin(n):
    st = "";
    i=n;
    while i > 0:
        if i == n:
            st = "sin("+str(i)+")"
        else:
            symb = "-" if i%2 else "+"
            st = "sin("+str(i)+symb+st+")"
        i -= 1
    #sin(1–sin(2+sin(3 - sin(4)))
    #sin(1–sin(2+sin(3))
    #sin(1-sin(2))
    #sin(1)
    return st

def funSin2(n):
    st = "";
    i=0;
    while i<n:
        if (i == 0):
            st = funSin(i+1)+"+"+str(n)
        else:
            st = "("+st+")"+funSin(i+1)+"+"+str(n-i)
        i += 1
    return st;
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
k = funSin2(count)
print(k)
