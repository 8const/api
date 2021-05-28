package handlers

import (
    "api/internal/service/handlers/pg"
    "context"
    "strconv"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "database/sql"
    "gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
    logCtxKey ctxKey = iota
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
    return func(ctx context.Context) context.Context {
        return context.WithValue(ctx, logCtxKey, entry)
    }
}

func Log(r *http.Request) *logan.Entry {
    return r.Context().Value(logCtxKey).(*logan.Entry)
}


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

//a closure to pass *sql.db to handler
//handler for search by id
func Ser(db *sql.DB) http.HandlerFunc {
    fn := func(w http.ResponseWriter, r *http.Request) {

        //get id from request
        i, err := strconv.Atoi(r.FormValue("id"))
        if err != nil {
            panic(err)
        }

        //query for row with such id
/*        err = db.QueryRow("SELECT blob FROM blobs WHERE id=($1);", i).Scan(&bb)
        if err != nil {
            panic(err)
        }
*/

        q := pg.NewBlobsQ(db)
        bb := q.SelectById(i)

        //decode row to struct 
        b := Dunmarshal(bb)

        //fill json with struct data
        id   := strconv.Itoa(int(b.User_id))
        name := string(b.User_name)
        json := []byte("{'user_id':'" + id + "' ,'user_name':'" + name + "'}")
        w.Header().Set("Content-Type", "application/json")

        //send it
        w.Write(json)
    }
    return http.HandlerFunc(fn)
}

func Lis(db *sql.DB) http.HandlerFunc {
    fn := func(w http.ResponseWriter, r *http.Request) {

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
    }

    return http.HandlerFunc(fn)
}



func Rem(db *sql.DB) http.HandlerFunc {
    fn := func(w http.ResponseWriter, r *http.Request) {

        id, err := ioutil.ReadAll(r.Body)
        if err != nil {
            panic (err)
        }

        _, err = db.Exec("DELETE FROM blobs WHERE id=" + string(id) + ";")
        if err != nil {
            panic(err)
        }
    }
    return http.HandlerFunc(fn)
}


func New(db *sql.DB) http.HandlerFunc {
    fn := func(w http.ResponseWriter, r *http.Request) {
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
    }
    return http.HandlerFunc(fn)
}
