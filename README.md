CLONE IT:

      $ cd $GOPATH/src
      $ git clone git@github.com:8const/api.git
      $ cd api/

SETTING UP DB:

      $ docker run --name postgresql-container -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres 

SET UP ENVIRONMENT:

      $ export GO111MODULE=on
      $ export KV_VIPER_FILE=config.yaml
      $ go mod init
      $ go mod vendor

BUILD IT:

      $ go build -o api main.go 

APPLY MIGRATIONS:

      $ ./api migrate up


RUN IT IN THE BACKGROUND:
      
      $ ./api run service &
      
      
USE PYTHON CLIENT:

      $ python3
      >>> from client import *
      /*help message will be printed*/
