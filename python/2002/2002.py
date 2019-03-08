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

users = {}
count = int(s[0])
i = 1
while i <= count:
    s.append(input())
    params = s[i].split(" ")
    action = params[0]
    if action == "register":
        login = params[1]
        pas = params[2]
        if login in users:
            print("fail: user already exists")
        else:
            users[login] = {'login':login, 'pass':pas, 'status':'off'}
            print('success: new user added')
    elif action == "login":
        login = params[1]
        pas = params[2]
        if login in users:
            if users[login]['pass'] == pas:
                if users[login]['status'] == 'off':
                    users[login]['status'] = 'on'
                    print('success: user logged in')
                else:
                    print('fail: already logged in')
            else:
                print('fail: incorrect password')
        else:
            print('fail: no such user')
    elif action == "logout":
        login = params[1]
        if login in users:
            if users[login]['status'] == 'on':
                users[login]['status'] = "off"
                print("success: user logged out")
            else:
                print('fail: already logged out')
        else:
            print('fail: no such user')
    i += 1
