package handlers

import (
	"database/sql/driver"
	"encoding/json"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

//driverValue - converts interface into db supported type
func DriverValue(data interface{}) (driver.Value, error) {
	data, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New("failed to marshal details")
	}

	return data, nil
}

//driveScan - converts jsonb into type struct
func DriveScan(src, dest interface{}) error {
	data, err := ConvertJSONB(src)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, dest)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal jsonb")
	}

	return nil
}

func ConvertJSONB(src interface{}) ([]byte, error) {
	var data []byte
	switch rawData := src.(type) {
	case []byte:
		data = rawData
	case string:
		data = []byte(rawData)
	default:
		return nil, errors.New("Unexpected type for jsonb")
	}

	return data, nil
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

