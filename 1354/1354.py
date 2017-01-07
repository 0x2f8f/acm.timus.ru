import pprint, sys, math, hashlib

def printPolindrom(word, i, k):
    j = i-k
    word2 = ''
    while j>=0:
        word2+=word[j]
        j-=1
    print(word+word2)
    sys.exit()
"""
fin = open("e.in")
s = fin.read().split('\n')
#s = fin.read()
#s = fin.readline().strip()
word = s[0]
"""
word = input()

i = 0
#d1 = {}
#d2 = {}
k = 0
n = len(word)
r = int(n/2)

if n==1:
    print(word+word)
    sys.exit()

while i < n:
    #d2[i] = 0
    k=0
    #while (i-d2[i]-1) >= 0 and (i+d2[i]) < n and word[i-d2[i]-1] == word[i+d2[i]]:
    while (i-k-1) >= 0 and (i+k) < n and word[i-k-1] == word[i+k]:
        #d2[i]+=1
        k+=1
    if k+i==n and (k<r or r==1 and n>2):
        #print('@@@-'+word[i]+'---'+str(k))
        printPolindrom(word, i-1, k)
    #d1[i] = 1
    k = 1
    #while (i-d1[i]) >= 0 and (i+d1[i]) < n and word[i-d1[i]] == word[i+d1[i]]:
    while (i-k) >= 0 and (i+k) < n and word[i-k] == word[i+k]:
        #d1[i]+=1
        k+=1
    if k+i==n and k!=1 and k<r+1:
        #print('!!!-'+word[i]+'---'+str(k))
        #printPolindrom(word, i-1 if k==1 else i, k)
        printPolindrom(word, i, k)
    i+=1
#print('END')
printPolindrom(word, n-1, 1)
#print(d1, d2)
