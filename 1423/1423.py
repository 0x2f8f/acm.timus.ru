import pprint, sys, math

"""
f = open('a.in', 'w')
f.write("250000")
f.write('\n')

i = 0
while i<(250000-1):
    f.write('a')
    i+=1
f.write('a')
f.write('\n')

i = 0
while i<(250000-1):
    f.write('a')
    i+=1
f.write('b')

f.close()

sys.exit()
"""
fin = open("a.in")
s = fin.read().split('\n')
#s = fin.read()
#s = fin.readline().strip()
"""
s = []
s.append(input())
s.append(input())
s.append(input())
"""

"""
s = sys.stdin
"""

#pprint.pprint(s)
#sys.exit()


count = int(s[0])
word = s[1]
word2 = s[2]

if word == word2:
    print(0)
    sys.exit()

l1 = len(word)
l2 = len(word2)

#считаем все степени p
p  = 2
pPow = {0:1}
i = 1
while i < l1:
    pPow[i] = pPow[i-1]*p
    i += 1

#считаем хеши префиксов
# 1ое слово
i = 0
h1 = {}
hc = {}
aOrd = ord('!')
while i < l1:
    h1[i] = (ord(word[i]) - aOrd + 1)*pPow[i]
    if i:
        h1[i] += h1[i-1]
    hc[h1[i]]=i
    i += 1
    
# 2ое слово
i = 0
h2 = {}
while i < l2:
    h2[i] = (ord(word2[i]) - aOrd + 1)*pPow[i]
    if i:
        h2[i] += h2[i-1]
    i += 1

h3 = {}
i = l2-1
j = 0
while i>=0:
    h3[i] = ord(word2[i]) - aOrd + 1
    if j:
        h3[i] += h3[i+1]*p
    if h3[i] in hc:
        #print( word2[i:l2]+word2[0:i])
        #проверяем всю оставшуюся часть
        i1 = j+1
        i2 = 0
        s1 = h1[i1+i-1] - h1[i1-1]
        s2 = h2[0+i-1]
        if s1 == (s2 * pPow[i1-i2]):
        #if  word2[i:l2]+word2[0:i] == word:
            print(i)
            sys.exit()
    i -= 1
    j += 1
print("-1")

#pprint.pprint(h)
#pprint.pprint(h2)

"""
i = l2 - 1
st = ""
while i >= 0:
    st = word2[i] + st    
    if word.rfind(st) > -1 and word2[i:l2]+word2[0:i] == word:
        print(i)
        sys.exit()
    i -= 1
print("-1")
"""
