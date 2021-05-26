Before running do somthing like $ go mod vendor to collect the dependencies


Run api with:


$ ./run



It assumes a PostgresSQL db is running on localhost with user/password/etc just like in internal/service/router.go at line ~50  
do smth like $ docker load with db.tar to get it

there'll be a docker run command here for it


How to use the client:


$ python3
>>> from client.py import *
>>> /*call some stuff */


A lot of files are not used; Some important stuff is in internal/service/router.go
