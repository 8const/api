package handlers
import(
    "github.com/go-chi/chi"
	"encoding/json"
    "net/http"
    "api/internal/service/handlers/pg"
    "database/sql"
)

//handler for search by id
func Ser(db *sql.DB) http.HandlerFunc {
    fn := func(w http.ResponseWriter, r *http.Request) {

        //get id from request
        id := chi.URLParam(r, "id")

		if ! IsInt(id) {
			http.Error(w, "Bad Request (id has to be an integer)", 400)
            return
		}

        q := pg.NewBlobsQ(db)
        bb, err := q.SelectById(string(id))
        if err != nil {
			http.Error(w, "Not Found (no such blob in db)", 404)
            return
        }

        //decode row to struct 
        var b Dblob
        err = DriveScan(bb, &b)
        if err != nil {
			http.Error(w, "Internal Server Error", 500)
			return
        }

        res := SearchResult{b}
        json, err := json.Marshal(res)
        if err != nil {
			http.Error(w, "Internal Server Error", 500)
			return
        }

        w.WriteHeader(200)
        w.Write(json)
    }
    return http.HandlerFunc(fn)
}
