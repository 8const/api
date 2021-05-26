Run api with:


$ ./run



It assumes a lot of dependencies and most importantly a PostgresSQL db running on localhost with config: postgres://postgres:postgres@127.0.0.1:5432/stuff?sslmode=disable.
DB has to have a table called blobs described below.

id |                   blob                       
----+-------------------------------------------                                                     
3 | {"User_id": 3, "User_name": "lol"}         
 
 
 
                                                Table "public.blobs"                                  
 Column |  Type   | Collation | Nullable |              Default              | Storage  | Stats target | Description   
--------+---------+-----------+----------+-----------------------------------+----------+--------------+------------- 
 id     | integer |           | not null | nextval('blobs_id_seq'::regclass) | plain    |              |                 
 blob   | jsonb   |           |          |                                   | extended |              |                


How to use the client:


$ python3
>>> from client.py import *
>>> /*call some stuff */

Some files are not really used; Some important stuff is in internal/service/router.go
