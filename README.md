SETTING UP DB:

      $ ./db.sh
      $ docker exec -it postgres bash
      container/# createdb -h localhost -p 5432 -U postgres stuff 
      container/# psql -U postgres
      postgres=#  \c stuff 
      postgres=#  CREATE TABLE blobs (id serial PRIMARY KEY, blob jsonb);


RUN IT:

      $ ./run

CLIENT USAGE:

      $ python3
      >>> from client.py import *

Some files are not really used; Some important stuff is in internal/service/router.go
