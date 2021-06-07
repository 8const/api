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
    blob = '{"data":{"user_id":' + str(uid) + ', "user_name":"' + str(uname) + '"}}'
    print(requests.post("http://127.0.0.1:8000/blobs/new", data=blob))


def rem(id):
    print(requests.delete("http://127.0.0.1:8000/blobs/delete/"+str(id)))

def lis():
    print(requests.get("http://127.0.0.1:8000/blobs/all"))
    print(requests.get("http://127.0.0.1:8000/blobs/all").text)


def ser(id):
    print(requests.get("http://127.0.0.1:8000/blobs/search/"+str(id)))
    print(requests.get("http://127.0.0.1:8000/blobs/search/"+str(id)).text)
