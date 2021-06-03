package handlers
import (
    "database/sql"
    "net/http"
    "github.com/go-chi/chi"
    "api/internal/service/handlers/pg"
)

func Rem(db *sql.DB) http.HandlerFunc {
    fn := func(w http.ResponseWriter, r *http.Request) {

        //get id from request
        id := (chi.URLParam(r, "id"))

        //query db
        q := pg.NewBlobsQ(db)
        q.DeleteById(string(id))

        w.WriteHeader(201)

    }
    return http.HandlerFunc(fn)
}












