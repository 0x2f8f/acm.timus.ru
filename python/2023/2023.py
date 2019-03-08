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

users = {"Alice":1,
"Ariel":1,
"Aurora":1,
"Phil":1,
"Peter":1,
"Olaf":1,
"Phoebus":1,
"Ralph":1,
"Robin":1,
"Bambi":2,
"Belle":2,
"Bolt":2,
"Mulan":2,
"Mowgli":2,
"Mickey":2,
"Silver":2,
"Simba":2,
"Stitch":2,
"Dumbo":3,
"Genie":3,
"Jiminy":3,
"Kuzko":3,
"Kida":3,
"Kenai":3,
"Tarzan":3,
"Tiana":3,
"Winnie":3}

status = 1
summ = 0
count = int(s[0])
i = 1
while i <= count:
    s.append(input())
    step = users[s[i]]
    if step > status:
        summ += step-status
        status = step
    elif step < status:
        summ += status-step
        status = step
    i += 1
print(summ)
