package handlers
import (
    "encoding/json"
    "io/ioutil"
    "database/sql"
    "net/http"
    "api/internal/service/handlers/pg"
    "api/internal/service/handlers/models"
)

func New(db *sql.DB) http.HandlerFunc {
    fn := func(w http.ResponseWriter, r *http.Request) {
        //get request body
        body, err := ioutil.ReadAll(r.Body)
        if err != nil {
			http.Error(w, "Bad request (please provide a valid body)", 400)
            return
        }

        //put json data into struct
        var b models.DataDblob
        err = json.Unmarshal(body, &b)
        if err != nil {
			http.Error(w, "Bad request (please provide a valid body)", 400)
            return
        }

        //struct to DB's json
        bb, err := models.DriverValue(b.Data)
        if err != nil {
			http.Error(w, "Internal Server Error", 500)
            return
        }

        //actually always ok
        if newBlob, ok := bb.([]byte); ok {
            q := pg.NewBlobsQ(db)
            q.InsertBlob(newBlob)
        } else {
			http.Error(w, "Internal server error", 500)
            return
        }

    w.WriteHeader(201)
    }

    return http.HandlerFunc(fn)
}
