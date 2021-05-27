package service

import (
    "database/sql"
    "net/http"
    "github.com/go-chi/chi"
    "gitlab.com/distributed_lab/ape"
    "api/internal/service/handlers"
)



func (s *service) router() chi.Router {
    db, err := sql.Open("postgres",
    "postgres://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable")

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
    r.Get("/", handlers.Ser(db))

    //endpoint to list all rows
    r.Get("/all", handlers.Lis(db))

    //endpoint to create new rows 
    r.Post("/", handlers.New(db))

    //endpoint to delete row by id
    r.Delete("/", handlers.Rem(db))

    http.ListenAndServe(":8080", r)
    return r
}
