package json

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
)

type A struct {
	Id int64 `json:"aaa"`
}

type B map[string]json.RawMessage

func TestJsonUnmarshal(t *testing.T) {
	str := "{\"aaa\": 3548927381209002384}"
	aObj := &A{}
	bObj := &B{}

	_ = Unmarshal([]byte(str), aObj)
	fmt.Println(aObj)

	_ = Unmarshal([]byte(str), bObj)
	fmt.Printf("%T %v\n", (*bObj)["aaa"], (*bObj)["aaa"])
	fmt.Println(RawMsg((*bObj)["aaa"]).Int64())

	a, _ := RawMsg((*bObj)["aaa"]).Int64()
	(*bObj)["aaa"] = json.RawMessage(strconv.FormatInt(a+1, 10))

	b, _ := Marshal(bObj)
	fmt.Println(string(b))
}
