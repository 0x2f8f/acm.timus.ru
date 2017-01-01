import pprint
import sys

fin = open("a.in")
#s = fin.read()
s = fin.read().split('\n')
#s = fin.readline().strip()

pprint.pprint(s)
sys.exit()

#s = stdin.read()
vars = s[0]
#vars = input().split(' ')

a = int(vars[0])
b = int(vars[1])
result = a+b
print(result)
