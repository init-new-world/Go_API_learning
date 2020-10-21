import requests
import json

class Server:

    def __init__(self,addr="127.0.0.1",port=8086,header=dict()):
        self.addr = addr
        self.port = port
        self.header = header

    def url_create(self,url):
        return "http://" + self.addr + ':' + str(self.port) + url

def post_test(server,url='',data=dict(),get_req=False,is_json=False):
    r = None
    if is_json:
        r = requests.post(server.url_create(url),data=json.dumps(data),headers=server.header)
    else:
        r = requests.post(server.url_create(url),data=data,headers=server.header)
    if get_req == True:
        print(r.status_code)
        print(r.text)

def get_test(server,url='',data=dict(),get_req=False):
    r = requests.get(server.url_create(url),params=data,headers=server.header)
    if get_req == True:
        print(r.status_code)
        print(r.text)

def test_create():
    server = Server(header={'Content-Type': 'application/json'})
    user_data = dict()
    user_data['username'] = 'administrator'
    user_data['password'] = 'Hello,world!'
    post_test(server,"/v1/user/create",user_data,get_req=True,is_json=True)
    user_data['username'] = 'administrator'
    user_data['password'] = ''
    post_test(server,"/v1/user/create",user_data,get_req=True,is_json=True)
    user_data['username'] = ''
    user_data['password'] = 'Hello,world!'
    post_test(server,"/v1/user/create",user_data,get_req=True,is_json=True)

def test_monitor():
    server = Server()
    get_test(server,'/sd/monitor',get_req=True)

if __name__ == '__main__':
    test_monitor()
    test_create()
    test_monitor()
