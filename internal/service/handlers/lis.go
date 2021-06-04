package handlers
import (
    "fmt"
    "encoding/json"
    "net/http"
    "database/sql"
    "api/internal/service/handlers/pg"
)

func Lis(db *sql.DB) http.HandlerFunc {
    fn := func(w http.ResponseWriter, r *http.Request) {

        q := pg.NewBlobsQ(db)
        rows := q.SelectAll()

        var data []Row
        fmt.Println(rows)
        for rows.Next() {

            var bb []byte
            var b Dblob
            var i int64

            err := rows.Scan(&i, &bb)
            if err != nil {
                w.WriteHeader(500)
                return
            }

            err = DriveScan(bb, &b)
            if err != nil {
                w.WriteHeader(500)
                return
            }

            jblob := Jblob{b.User_id, b.User_name}

            row := Row{i, jblob}
            fmt.Println(row)
            data = append(data, row)
        }

        x := FinalResponse{data}
        y, err := json.Marshal(x)
        if err != nil {
            w.WriteHeader(500)
            return
        }

        w.Write(y)
    }

    return http.HandlerFunc(fn)
}


