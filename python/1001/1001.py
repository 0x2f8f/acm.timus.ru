import pprint, sys, math

"""
fin = open("a.in")
s = fin.read().split('\n')
#s = fin.read()
#s = fin.readline().strip()
"""

"""
s = []
s.append(input())
"""


s = sys.stdin

#pprint.pprint(s)
#sys.exit()

nums = []
i = 0
for line in s:
    for num in line.split():
        num = math.sqrt(int(num))
        nums.insert(i, num)
        i += 1
        
while i > 0:
    i -= 1
    print("%.4f" % (nums[i]))
