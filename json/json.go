package json

import (
	"encoding/json"
	"strconv"

	jsoniter "github.com/json-iterator/go"
)

var Jsoniter = jsoniter.Config{
	EscapeHTML:             false,
	SortMapKeys:            false,
	ValidateJsonRawMessage: true,
	UseNumber:              true,
}.Froze()

func Marshal(v interface{}) ([]byte, error) {
	return Jsoniter.Marshal(v)
}

func MarshalToString(v interface{}) (string, error) {
	return Jsoniter.MarshalToString(v)
}

func Unmarshal(data []byte, v interface{}) error {
	return Jsoniter.Unmarshal(data, v)
}

func UnmarshalFromString(data string, v interface{}) error {
	return Jsoniter.UnmarshalFromString(data, v)
}

type RawMsg json.RawMessage

func (n RawMsg) Float64() (float64, error) {
	return strconv.ParseFloat(string(n), 64)
}

func (n RawMsg) Int64() (int64, error) {
	return strconv.ParseInt(string(n), 10, 64)
}

func (n RawMsg) UInt64() (uint64, error) {
	return strconv.ParseUint(string(n), 10, 64)
}
