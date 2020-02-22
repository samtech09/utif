//Package utif stands for Utility Functions, providing different functions to facilitate common tasks
package utif

import (
	"database/sql"
	"encoding/json"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"
)

//StructToJSON convert given struct to JSON []byte
func StructToJSON(d interface{}) ([]byte, error) {
	data, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//StructToJSONString convert given struct to JSON
func StructToJSONString(d interface{}) (string, error) {
	data, err := json.Marshal(d)
	if err != nil {
		return "", err
	}
	var sb strings.Builder
	sb.Write(data)
	//jsonstring := string(data[:]) //BytesToString(data)
	return sb.String(), nil
}

//JSONStringToStruct convert JSON string to given struct
func JSONStringToStruct(jsonstring string, dest interface{}) error {
	if err := json.Unmarshal([]byte(jsonstring), dest); err != nil {

		return err

	}

	return nil
}

// //BytesToString convert []byte to string
// func BytesToString(b []byte) string {
// 	//bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
// 	//sh := reflect.StringHeader{bh.Data, bh.Len}
// 	//return *(*string)(unsafe.Pointer(&sh))

// 	return string(b[:])
// }

//ToString try to convert given interface to string
func ToString(t interface{}) (string, bool) {
	value, ok := t.(string)
	if ok {
		return value, true
	}
	if val, ok := t.([]byte); ok {
		return string(val), true
	}
	// return default value
	return "", false
}

//ToStringx try to convert given struct to string
// if any error occur then return blank string
func ToStringx(t interface{}) string {
	value, ok := t.(string)
	if ok {
		return value
	}
	if val, ok := t.([]byte); ok {
		return string(val)
	}
	// return default value
	return ""
}

// StoI converts slice of struct to slice of interface
func StoI(array interface{}) []interface{} {

	v := reflect.ValueOf(array)
	t := v.Type()

	if t.Kind() != reflect.Slice {
		log.Panicf("`array` should be %s but got %s", reflect.Slice, t.Kind())
	}

	result := make([]interface{}, v.Len(), v.Len())

	for i := 0; i < v.Len(); i++ {
		result[i] = v.Index(i).Interface()
	}

	return result
}

//TimeValue return time value from sql.NullTime, if has invalid value then 1900-01-01 will be returned.
func TimeValue(timeValue sql.NullTime) time.Time {
	if timeValue.Valid {
		return timeValue.Time
	}
	return time.Date(1901, 1, 1, 0, 0, 0, 0, time.Now().Local().Location())
}

//StringValue return string value from sql.NullString, if has invalid value then zero value will be returned.
func StringValue(stringValue sql.NullString) string {
	if stringValue.Valid {
		return stringValue.String
	}
	return ""
}

//IntValue return int value from sql.NullInt32, if has invalid value then zero value will be returned.
func IntValue(intValue sql.NullInt32) int32 {
	if intValue.Valid {
		return intValue.Int32
	}
	return 0
}

//FloatValue return float value from sql.NullFloat64, if has invalid value then zero value will be returned.
func FloatValue(intValue sql.NullFloat64) float64 {
	if intValue.Valid {
		return intValue.Float64
	}
	return 0
}

//IntSliceToCSV convert slice of int to comma separated string
func IntSliceToCSV(a []int, sep string) string {
	if len(a) == 0 {
		return ""
	}

	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.Itoa(v)
	}
	return strings.Join(b, sep)
}
