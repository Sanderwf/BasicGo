package convert

import (
	"testing"
)

func TestStruct2JsonString(t *testing.T) {
	apiUserIp := apiUserIps{
		[]userIps{
			{
				"lijin",
				[]ipStatus{
					{
						"192.168.0.1",
						-1,
					},
				},
			},
			{
				"wangfei",
				[]ipStatus{
					{
						"192.168.0.1",
						-1,
					},
					{
						"192.168.0.2",
						0,
					},
				},
			},
		},
	}
	t.Log("struct -> JsonStr")
	jsonString, err := Struct2Jsonstr(apiUserIp)
	if err != nil {
		t.Log("\n err: ", err)
	}
	if len(jsonString) > 0 {
		t.Log("\n", jsonString)
	}
}

func TestJsonString2Struct(t *testing.T) {
	str := "{\"UserIps\":[{\"UserId\":\"lijin\",\"Ips\":[{\"Ip\":\"192.168.0.1\",\"IsEffective\":-1}]}," +
		"{\"UserId\":\"wangfei\",\"Ips\":[{\"Ip\":\"192.168.0.1\",\"IsEffective\":-1},{\"Ip\":\"192.168.0.2\",\"IsEffective\":0}]}]}\n"

	//str1:= "{\"userIpSegs\":{\"segments\":[\"192.168.0.10-192.168.0.20\"],\"userIps\":[{\"ip\":\"192.168.0.1\",\"message\":\"123\"}]}}"

	apiUserIps, err := Jsonstr2Structstr(str)
	if err != nil {
		t.Log("\n err: ", err)
	} else {
		t.Log("\n", apiUserIps)
	}
}

// 拷贝JSON串到Goland，可以直接生成对应的结构体构造
type T struct {
	UserIps []struct {
		UserId string `json:"UserId"`
		Ips    []struct {
			Ip          string `json:"Ip"`
			IsEffective int    `json:"IsEffective"`
		} `json:"Ips"`
	} `json:"UserIps"`
}
