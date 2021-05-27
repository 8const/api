package service

import (
    "strconv"
    "encoding/json"
    "database/sql"
    "io/ioutil"
    "net/http"
    "github.com/go-chi/chi"
    "gitlab.com/distributed_lab/ape"
    "api/internal/service/handlers"
)


//generally b is used for blob as a a struct
//bb is bytes of marshalled struct

//for encoding/decoding incoming client's json
type Jblob struct {
    User_id   int64  `json:'user_id'`
    User_name string `json:'user_name'`
}

//for encoding/decoding DB's jsonb
type Dblob struct {
    User_id   int64  `db:'user_id'`
    User_name string `db:'user_name'`
}


//request json to struct
func Junmarshal(bb []byte) Jblob {
    var res Jblob
    err := json.Unmarshal(bb, &res)
    if err != nil {
        panic(err)
    }
    return res
}

//struct to DB's json
func Dmarshal(b Dblob) []byte{
    res, err := json.Marshal(b)
    if err != nil {
        panic(err)
    }
    return res
}

//DB's json to struct
func Dunmarshal(bb []byte) Dblob {
    var res Dblob
    err := json.Unmarshal(bb, &res)
    if err != nil {
        panic(err)
    }
    return res
}

//used to make string for json response nicer by removing whitespace and extra ','
func remel(slice []rune, i int) []rune {
    return append(slice[:i], slice[i+1:]...)
}


func (s *service) router() chi.Router {
    db, err := sql.Open("postgres",
    "postgres://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable")

    err = db.Ping()
    if err != nil {
        print("\nNO SUCH DB RUNNING OR WRONG AUTH\n")
    } else {
        print("\nNICE: DB CONNECTED\n")
    }

    r := chi.NewRouter()

    r.Use(
          ape.RecoverMiddleware(s.log),
          ape.LoganMiddleware(s.log),
          // this line may cause compilation error 
          //but in general case `dep ensure -v` will fix it
          ape.CtxMiddleware(handlers.CtxLog(s.log),
          ),
    )

    //endpoint for search by id
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {

        //get id from request
        i, err := strconv.Atoi(r.FormValue("id"))
        if err != nil {
            panic(err)
        }

        //query for row with such id
        var bb []byte
        err = db.QueryRow("SELECT blob FROM blobs WHERE id=($1);", i).Scan(&bb)
        if err != nil {
            panic(err)
        }

        //decode row to struct 
        b := Dunmarshal(bb)

        //fill json with struct data
        id   := strconv.Itoa(int(b.User_id))
        name := string(b.User_name)
        json := []byte("{'user_id':'" + id + "' ,'user_name':'" + name + "'}")
        w.Header().Set("Content-Type", "application/json")

        //send it
        w.Write(json)
    })

    //endpoint lists all rows
    r.Get("/all", func(w http.ResponseWriter, r *http.Request) {

        //get all rows from db
        rows, err := db.Query("SELECT * FROM blobs")
        if err != nil {
            panic(err)
        }
        defer rows.Close()

        //form a string that looks like json
        //fill it with rows
        list := "["
        for rows.Next() {
            var b Dblob
            bb := Dmarshal(b)
            i := 0
            err = rows.Scan(&i, &bb)
            if err != nil {
            panic(err)
            }

            id   := strconv.Itoa(int(Dunmarshal(bb).User_id))
            name := string(Dunmarshal(bb).User_name)

            list += ("{'user_id':'" + id + "' ,'user_name':'" + name + "'}, ")
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
    })

    r.Post("/", func(w http.ResponseWriter, r *http.Request) {

        //get request body
        body, err := ioutil.ReadAll(r.Body)
        if err != nil {
            panic (err)
        }

        //json from req body to struct 
        b := Junmarshal(body)

        //struct to DB's json
        bb := Dmarshal(Dblob{b.User_id, b.User_name})

        //insert it into DB 
        _, err = db.Exec("INSERT INTO blobs (blob) VALUES ($1)", bb)
        if err != nil {
            panic(err)
        }
    })

    //delete by id endpoint
    r.Delete("/", func(w http.ResponseWriter, r *http.Request) {

        id, err := ioutil.ReadAll(r.Body)
        if err != nil {
            panic (err)
        }

        _, err = db.Exec("DELETE FROM blobs WHERE id=" + string(id) + ";")
        if err != nil {
            panic(err)
        }
    })

    http.ListenAndServe(":8080", r)
    return r
}
