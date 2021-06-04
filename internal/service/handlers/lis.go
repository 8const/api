package handlers
import (
    "fmt"
    "encoding/json"
//    "strconv"
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
       // list := "["
        var data []Row
        for rows.Next() {

            //bb := Dmarshal(b)
            var bb []byte
            var b Dblob
            //var x Row
            var i int64

            err := rows.Scan(&i, &bb)
            if err != nil {
                panic(err)
            }

            err = DriveScan(bb, &b)
            if err != nil {
                print("\n41\n")
            }

            jblob := Jblob{b.User_id, b.User_name}



            row := Row{i, jblob}
            fmt.Println(row)
            data = append(data, row)

            //here struct ntiaalised with i and bb but bb is now a struct for json
           // Row{i, b}

          //  fmt.Println(b)
/*
            DriveScan(bb, &b)
            if err != nil {
                w.WriteHeader(500)
                return
            }
            ///////////WORK HERE

            jb := Jblob{b.User_id, b.User_name}

            data = append(data, jb)

            list += ("{'user_id':'" + strconv.Itoa(int(b.User_id)) + "' ,'user_name':'" + b.User_name+ "'}, ")
            */
        }
        print("\n---------------------\n")
        fmt.Println(data)
        print("\n---------------------\n")
        //list += "]"


        x := FinalResponse{data}
        y, err := json.Marshal(x)
        if err != nil {
            print("AAAAAAAAAAA")
        }

        fmt.Println(y)



        //make it pretty
      //  sl := []rune(list)
       // sl = remel(sl, len(sl)-3)
       // sl = remel(sl, len(sl)-2)
       // w.Header().Set("Content-Type", "application/json")
        //send it
       /* data, err := json.Marshal(data)
        if err != nil {
            print("\nERROR\n")
        }
        */
        w.Write(y)
    }

    return http.HandlerFunc(fn)
}


