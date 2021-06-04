package pg

import (
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

func (q  *BlobsQ) SelectById(id string) ([]byte, error) {
    query, _, _ := q.sqlSelect.Where("id=$1").ToSql()
    var bb []byte
    err := ((q.db).QueryRow(query, id)).Scan(&bb)
    if err != nil {
        return nil, err
    }
    return bb, nil
}

func (q  *BlobsQ) SelectAll() *sql.Rows {
    query, _, _ := q.sqlSelectStar.ToSql()
    rows, err := ((q.db).Query(query))
    if err != nil {
        panic(err)
    }
    return rows
}

func (q  *BlobsQ) DeleteById(id string) (int, error) {
    query, _, _ := q.sqlDelete.Where("id=$1").ToSql()
    res, err := q.db.Exec(query, id)
    if err != nil {
        return -1, err
    }

    nChanges, err := res.RowsAffected()
    if err != nil {
        return -1, err
    } else {
        return int(nChanges), nil
    }
}

func (q  *BlobsQ) InsertBlob(blob []byte) {
    query, _, _ := q.sqlInsert.Values("($1)").PlaceholderFormat(sq.Dollar).ToSql()
    _, err := (q.db).Exec(query, blob)
    if err != nil {
        panic(err)
    }
}
