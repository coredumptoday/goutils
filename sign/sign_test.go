package sign

import (
	"crypto"
	"fmt"
	"net/url"
	"testing"
)

//一次性输入进行hash计算
func TestMds(t *testing.T) {
	ob := mds("abcdefghijklmn")
	fmt.Println(ob.MD5().ToHex().ToString())                //计算md5
	fmt.Println(ob.SHA1().ToHex().MD5().ToHex().ToString()) //计算md5(sha1(str))
}

//多次拼接数据进行hash计算
func TestSHA1Sign(t *testing.T) {
	h := NewSHA1()
	h.WriteString("abc")
	h.WriteString("def")
	h.WriteString("hig")
	h.WriteString("klm")

	sha1Sum, err := h.Sum()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(sha1Sum.ToHex().ToString()) //sha1(abcdefhigklm)
}

/**
 * URL请求数据签名
 * 请求参数为 aaa=111&bbb=222&ccc=333&ddd=444&fff=555
 * 签名规则为 md5( appkey + 请求参数看key降序排序，kv之间=分割，参数之间&分割 )
 */
func TestMD5SignFromQuery(t *testing.T) {
	queryStr := "aaa=111&bbb=222&ccc=333&ddd=444&fff=555"
	q, _ := url.ParseQuery(queryStr)

	builder := NewSignBuilderWithQuery(crypto.MD5, q)
	builder.DESCSort()
	builder.WritePrefixString("appkey")
	builder.SetKVSepStr("=").SetKVParamByte([]byte("&"))
	md5Sum, err := builder.Sign()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(md5Sum.ToHex().ToString())
}

/**
 * 根据map[string]string 进行签名
 * 签名规则为 md5( appkey + aaa=111&bbb=222&ccc=333&ddd=444&fff=555 + appkey )
 */
func TestMD5SignFromMap(t *testing.T) {
	m := map[string]string{
		"aaa": "111",
		"bbb": "222",
		"ccc": "333",
		"ddd": "444",
	}

	builder := NewSignBuilderWithMap(crypto.MD5, m)
	builder.Set("fff", "555")
	builder.ASCSort()
	builder.WritePrefixString("appkey")
	builder.SetKVSepStr("=").SetKVParamByte([]byte("&"))
	md5Sum, err := builder.SignWithPostfixStr("appkey")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(md5Sum.ToHex().ToString())
}
