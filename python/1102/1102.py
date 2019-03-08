import pprint, sys, math

def replaceWords(st):
    replaceWord = ["inputone", "outputone", "puton", "one", "output", "input", "out", "in"]
    for word in replaceWord:
        st = st.replace(word,"")
    return st

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

count = int(s[0])
i = 1
while i <= count:
    sr = input()
    word = replaceWords(sr)
    print("NO" if len(word) else "YES")
    i += 1
