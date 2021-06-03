package handlers
import (
    "strconv"
    "net/http"
    "database/sql"
    "api/internal/service/handlers/pg"
)
func Lis(db *sql.DB) http.HandlerFunc {
    fn := func(w http.ResponseWriter, r *http.Request) {

        //get all rows from db
        //rows, err := db.Query("SELECT * FROM blobs")


        q := pg.NewBlobsQ(db)
        rows := q.SelectAll()


        //form a string that looks like json
        //fill it with rows
        list := "["
        for rows.Next() {

            //bb := Dmarshal(b)
            var bb []byte
            var b Dblob
            i := 0

            err := rows.Scan(&i, &bb)
            if err != nil {
                panic(err)
            }

            DriveScan(bb, &b)
            if err != nil {
                w.WriteHeader(500)
                return
            }

            list += ("{'user_id':'" + strconv.Itoa(int(b.User_id)) + "' ,'user_name':'" + b.User_name+ "'}, ")
        }
        list += "]"

        //make it pretty
        sl := []rune(list)
        sl = remel(sl, len(sl)-3)
        sl = remel(sl, len(sl)-2)
        w.Header().Set("Content-Type", "application/json")
        json := []byte(string(sl))
        //send it
        w.Write(json)
    }

    return http.HandlerFunc(fn)
}


