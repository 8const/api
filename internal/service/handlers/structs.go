package handlers

import (
	"database/sql/driver"
    "unicode"
	"encoding/json"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func IsInt(s string) bool {
    for _, c := range s {
        if !unicode.IsDigit(c) {
            return false
        }
    }
    return true
}



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

//generally b is used for blob as a struct
//bb is bytes of marshalled struct

//for encoding/decoding incoming client's json
type Jblob struct {
    User_id   int64  `json:"user_id"`
    User_name string `json:"user_name"`
}

//for encoding/decoding DB's jsonb
type Dblob struct {
    User_id   int64  `db:"user_id"`
    User_name string `db:"user_name"`
}


type Row struct {
    Id int64   `json:"id"`
    Blob Jblob `json:"data"`
}


//to marshal list all
type FinalResponse struct {
    Data []Row `json:"data"`
}

//to marshal search by id results
type SearchResult struct {
    Data Dblob `json:"data"`
}

