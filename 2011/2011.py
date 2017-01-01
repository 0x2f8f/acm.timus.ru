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
s.append(input())

#pprint.pprint(s)
#sys.exit()


abc = {'1': 0, '2':0, '3':0}
count = int(s[0])
vars = s[1].split(' ')
i = 0
while i < count:
    abc[vars[i]]+=1
    i+=1

#сортируем
b = list(abc.items())
b.sort(key=lambda item: item[1])
abc = list(b)

summ = 0
denominator = 1
key = 0
while key < 3:
    abc[key] = list(abc[key])
    if abc[key][1] > 5:
        abc[key][1] = 5
    if abc[key][1]+summ > 6:
        abc[key][1] = 6 - summ
    summ += abc[key][1]
    factorial = math.factorial(abc[key][1])
    if factorial:
        denominator *= math.factorial(abc[key][1])
    key +=1
result = math.factorial(summ)/denominator

if result >= 6:
    print('Yes')
else:
    print('No')
