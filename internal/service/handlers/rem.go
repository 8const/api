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
        nChanges, err := q.DeleteById(string(id))
        if err != nil {
            w.WriteHeader(500)
            return
        }

        if nChanges == 0 {
            w.WriteHeader(404)
            return
        }

        w.WriteHeader(200)
    }
    return http.HandlerFunc(fn)
}












