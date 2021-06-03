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

        q := pg.NewBlobsQ(db)
        bb := q.SelectById(string(id))

        //decode row to struct 
        var b Dblob
        err := DriveScan(bb, &b)
        if err != nil {
            w.WriteHeader(500)
			return
        }

        json, err := json.Marshal(b)
        if err != nil {
            w.WriteHeader(500)
			return
        }

        w.WriteHeader(200)
        w.Write(json)
    }
    return http.HandlerFunc(fn)
}


