package service

import (
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
    "database/sql"
    "net/http"
    "github.com/go-chi/chi"
    "gitlab.com/distributed_lab/ape"
    "api/internal/service/handlers"
)
//struct to parse config.yaml
type con struct {
    LOG struct {
        DISSENTRY bool `yaml:disable_sentry`
    }
    DB struct {
        URL string`yaml:url`
    }
    LISTENER struct {
        ADDR string `yaml:addr`
    }

    COP struct {
        DISABLED bool   `yaml:disabled`
        ENDPOINT string `yaml endpoint`
        UPSTREAM string `yaml upstream`
        SNAME    string `yaml: service_name`
        SPORT    string `yaml:service_port`
    }
}

func (c *con) getConf() *con {
    yamlFile, err := ioutil.ReadFile("config.yaml")
    if err != nil {
        log.Printf("yamlFile.Get err   #%v ", err)
    }
    err = yaml.Unmarshal(yamlFile, c)
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
    }
    return c
}


func (s *service) router() chi.Router {
	var c con
	c.getConf()

    //DB is from config.yaml
    db, err := sql.Open("postgres", c.DB.URL)

    err = db.Ping()
    if err != nil {
        print("\nNO SUCH DB RUNNING OR WRONG AUTH\n")
    } else {
        print("\nNICE: DB CONNECTED\n")
    }

    r := chi.NewRouter()

    r.Use(
          ape.RecoverMiddleware(s.log),
          ape.LoganMiddleware(s.log),
          // this line may cause compilation error 
          //but in general case `dep ensure -v` will fix it
          ape.CtxMiddleware(handlers.CtxLog(s.log),
          ),
    )

    //endpoint to searche by id
    r.Get("/blobs/search/{id}", handlers.Ser(db))

    //endpoint to list all rows
    r.Get("/blobs/all", handlers.Lis(db))

    //endpoint to create new rows 
    r.Post("/blobs/new", handlers.New(db))

    //endpoint to delete row by id
    r.Delete("/blobs/delete/{id}", handlers.Rem(db))

    http.ListenAndServe(":8000", r)
    return r
}
