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
i = 1
k = 2
while i <= count:
    s.append(input())
    st = str(s[i])
    if (len(st.split('+one'))>1):
        k += 2
    else:
        k +=1
    i += 1
if (k == 13):
    k = 14
print(k*100)
