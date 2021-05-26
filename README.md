SETTING UP DB

$ ./db.sh

$ docker exec -it postgres bash

container/# createdb -h localhost -p 5432 -U postgres stuff 

container/# psql -U postgres

postgres=#  \c stuff 

postgres=#  CREATE TABLE blobs (id serial PRIMARY KEY, blob jsonb);


Now you can run api with:

$ ./run

How to use the client:


$ python3

&#62;&#62;&#62; from client.py import *

Some files are not really used; Some important stuff is in internal/service/router.go
