Run api with:


$ ./run



It assumes a lot of dependencies and most importantly a PostgresSQL db running on localhost with config: postgres://postgres:postgres@127.0.0.1:5432/stuff?sslmode=disable.
DB has to have a table called blobs described below.

&nbsp;&nbsp;&nbsp;id |                   blob                       
----+-------------------------------------------                                                     
&nbsp;&nbsp;&nbsp;&nbsp;3 | {"User_id": 3, "User_name": "lol"}         
 
 
 
 id     type:  integer; nullable:  not null; default nextval('blobs_id_seq'::regclass)         
 blob   type:  jsonb;


How to use the client:


$ python3
">>>" from client.py import *
">>>" /*call some stuff */

Some files are not really used; Some important stuff is in internal/service/router.go
