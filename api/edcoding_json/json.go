package edcoding_json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Id      int      `json:"xxx"` // 注意看 edcoding_json 对应的 key
	Name    string   `json:"name"`
	Name_Cn string   `json:"name_cn,omitempty"` // 若为零值，则不解析
	Books   []string `json:"books"`
	tel     string   // 小写开头为私有不会解析
	Addr    string   `json:"-"` // - ： 效果和小写开头一样，忽略此字段
}

func JsonDemo(user User) {
	if bytes, err := json.Marshal(user); err == nil {
		fmt.Println(string(bytes))
	}

	fmt.Println("\n 格式化打印，omitempty（零值不解析）")
	if bytes, err := json.MarshalIndent(user, "", "\t"); err == nil {
		fmt.Println(string(bytes))
	}

	fmt.Println("\n 解码：")
	s := `{"Id":1,"Name":"WANGFEI","Books":["a","b"]}`
	if err := json.Unmarshal([]byte(s), &user); err == nil {
		fmt.Println(user)
	}

	fmt.Println("\n 验证字符串是否是正确的JSON格式:")
	fmt.Println(json.Valid([]byte(s)))

}

func EncodeDemo(users []User) {
	fmt.Println(users)
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "\t")
	// 将json结果输出到标准流
	err := encoder.Encode(users)
	if err != nil {
		fmt.Println(err)
	}

	var out bytes.Buffer
	encoder = json.NewEncoder(&out)
	// 格式化处理
	encoder.SetIndent("", "\t")
	err = encoder.Encode(users)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("out.String\n", out.String())
	}
}
