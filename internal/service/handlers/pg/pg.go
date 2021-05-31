package pg

import (
    "fmt"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
)

type BlobsQ struct {
	db           *sql.DB
	sqlSelect     sq.SelectBuilder
	sqlSelectStar sq.SelectBuilder
    sqlDelete     sq.DeleteBuilder
    sqlInsert     sq.InsertBuilder
}

func NewBlobsQ(db *sql.DB) *BlobsQ {
	return &BlobsQ {
		db:            db,
        sqlSelect:     sq.Select("blob").From("blobs"),
        sqlSelectStar: sq.Select("*").From("blobs"),
        sqlDelete:     sq.Delete("blob").From("blobs"),
        sqlInsert:     sq.Insert("").Into("blobs").Columns("blob"),
    }
}

func (q  *BlobsQ) SelectById(id int) []byte {
    query, _, _ := q.sqlSelect.Where("id=$1").ToSql()
    fmt.Println(query)
    var bb []byte
    err := ((q.db).QueryRow(query, id)).Scan(&bb)
    if err != nil {
        panic(err)
    }
    return bb
}

func (q  *BlobsQ) SelectAll() *sql.Rows {
    query, _, _ := q.sqlSelectStar.ToSql()
    fmt.Println(query)
    rows, err := ((q.db).Query(query))
    if err != nil {
        panic(err)
    }
    return rows
}

func (q  *BlobsQ) DeleteById(id string) {
    query, _, _ := q.sqlDelete.Where("id=$1").ToSql()
    fmt.Println(query)
    _, err := ((q.db).Exec(query, id))
    if err != nil {
        panic(err)
    }
}

func (q  *BlobsQ) InsertBlob(blob []byte) {
    query, _, _ := q.sqlInsert.Values("($1)").PlaceholderFormat(sq.Dollar).ToSql()
    fmt.Println(string(query))
    _, err := (q.db).Exec(query, blob)
    if err != nil {
        panic(err)
    }
}
