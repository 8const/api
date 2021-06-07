package handlers
import (
	//"github.com/google/jsonapi"
    "database/sql"
    "net/http"
    "github.com/go-chi/chi"
    "api/internal/service/handlers/pg"
)

func Rem(db *sql.DB) http.HandlerFunc {
    fn := func(w http.ResponseWriter, r *http.Request) {

        //get id from request
        id := (chi.URLParam(r, "id"))

		if ! IsInt(id) {
			http.Error(w, "Bad Request (id has to be an integer)", 400)
            return
		}

        //query db
        q := pg.NewBlobsQ(db)
        nChanges, err := q.DeleteById(string(id))
        if err != nil {
			http.Error(w, "Internal server error", 500)
            return
        }

        if nChanges == 0 {
			http.Error(w, "Not Found (no such blob in db)", 404)
            return
        }

        w.WriteHeader(200)
    }
    return http.HandlerFunc(fn)
}












