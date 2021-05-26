import requests
print(
      "Assumes python in interactive shell mode; Usage:\n" 
      "In [1]: from client import * \n" 
      "You can now call:\n\n"
      "new(user_id, user_name)   to create a new row\n"
      "rem(id)                   to remove a row by id, where id is DB's id column\n"
      "lis()                     to list all rows\n"
      "ser(id)                   to search a row by id, where id is DB's id column\n"
     )

def new(uid, uname):
    blob = '{"user_id":' + str(uid) + ', "user_name":"' + str(uname) + '"}'
    requests.post("http://127.0.0.1:8080", data=blob)
    print(blob)


def rem(id):
    requests.delete("http://127.0.0.1:8080", data=str(id))

def lis():
    print(requests.get("http://127.0.0.1:8080/all").text)


def ser(id):
    print(requests.get("http://127.0.0.1:8080/?id="+str(id)).text)
