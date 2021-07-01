package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type A struct {
	Id int64 `json:"aaa"`
}

type B map[string]interface{}

func TestJsonUnmarshal(t *testing.T) {
	str := "{\"aaa\": 3548927381209002384}"
	aObj := &A{}
	bObj := &B{}

	_ = Unmarshal([]byte(str), aObj)
	fmt.Println(aObj)

	_ = Unmarshal([]byte(str), bObj)
	fmt.Printf("%T %v\n", (*bObj)["aaa"], (*bObj)["aaa"])
	fmt.Println((*bObj)["aaa"].(json.Number).Int64())
	fmt.Println(bObj)
}
