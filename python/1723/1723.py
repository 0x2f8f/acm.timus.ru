import pprint, sys, math

"""
fin = open("a.in")
s = fin.read().split('\n')
#s = fin.read()
#s = fin.readline().strip()
"""


s = []
s.append(input())


"""
s = sys.stdin
"""

#pprint.pprint(s)
#sys.exit()

st = s[0]
i = 0
words = {}
maxlen = 0
let = ""
while i<(len(st)):
    substr = st[i]
    #print(substr)
    i += 1
    if substr in words:
        words[substr] += 1
    else:
        words[substr]=1
    if maxlen < words[substr]:
        maxlen = words[substr]
        let = substr
print(let)
