package edcoding_json

import "testing"

var users = []User{
	{
		Id:      1,
		Name:    "WANGFEI",
		Name_Cn: "1",
		Books:   []string{"a1", "b1"},
		tel:     "185519941531",
		Addr:    "南京市建邺区1",
	},
	{
		Id:      2,
		Name:    "WANGFEI2",
		Name_Cn: "2",
		Books:   []string{"a2", "b2"},
		tel:     "185519941532",
		Addr:    "南京市建邺区2",
	},
}

var user = User{
	Id:      1,
	Name:    "WANGFEI",
	Name_Cn: "",
	Books:   []string{"a", "b"},
	tel:     "18551994153",
	Addr:    "南京市建邺区",
}

func TestJsonDemo(t *testing.T) {
	JsonDemo(user)
}

func TestEncodeDemo(t *testing.T) {
	EncodeDemo(users)
}
