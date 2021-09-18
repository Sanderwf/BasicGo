package encoding_base64

import (
	"encoding/base64"
	"fmt"
)

func Base64(){
	fmt.Println("实际应用：简单编码，字节切片持久化、网络传输")

	fmt.Println(base64.StdEncoding.EncodeToString([]byte("我是王飞")))
	decode, _ := base64.StdEncoding.DecodeString("5oiR5piv546L6aOe")
	fmt.Println(decode)


	fmt.Println(base64.RawURLEncoding.EncodeToString([]byte("我是王飞")))
	decode2, _ := base64.RawURLEncoding.DecodeString("5oiR5piv546L6aOe")
	fmt.Println(decode2)
}