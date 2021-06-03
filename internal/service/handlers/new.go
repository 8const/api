package handlers
import (
    "encoding/json"
    "io/ioutil"
    "database/sql"
    "net/http"
    "api/internal/service/handlers/pg"

)
func New(db *sql.DB) http.HandlerFunc {
    fn := func(w http.ResponseWriter, r *http.Request) {
        //get request body
        body, err := ioutil.ReadAll(r.Body)
        if err != nil {
            w.WriteHeader(400)
            return
        }

        //json from req body to struct 
        //b := Junmarshal(body)
        var b Jblob
        err = json.Unmarshal(body, &b)
        if err != nil {
            w.WriteHeader(400)
            return
        }


        //struct to DB's json
        //bb := Dmarshal(Dblob{b.User_id, b.User_name})
        bb, err := DriverValue(b)
        if err != nil {
            w.WriteHeader(400)
            return
        }
        if bbb, ok := bb.([]byte); ok {
            q := pg.NewBlobsQ(db)
            q.InsertBlob(bbb)
        } else {
            w.WriteHeader(500)
            return
        }

    w.WriteHeader(201)
    }

    return http.HandlerFunc(fn)
}
