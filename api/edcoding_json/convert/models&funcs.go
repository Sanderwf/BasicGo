package convert

import (
	"encoding/json"
)

type apiUserIps struct {
	UserIps []userIps
}

type userIps struct {
	UserId string
	Ips    []ipStatus
}

type ipStatus struct {
	Ip          string
	IsEffective int
}

func Struct2Jsonstr(input interface{}) (string, error) {
	bytes, err := json.Marshal(input)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func Jsonstr2Structstr(input string) (apiUserIps, error) {
	var apiUserIps apiUserIps
	err := json.Unmarshal([]byte(input), &apiUserIps)
	return apiUserIps, err
}

func Jsonstr2Structstr2(input string) (UserIpList, error) {
	var userIpSegs UserIpList
	err := json.Unmarshal([]byte(input), &userIpSegs)
	return userIpSegs, err
}

type UserIpList struct {
	UserIpSegs struct {
		Segments []string `json:"segments"`
		UserIps  []struct {
			Ip      string `json:"ip"`
			Message string `json:"message"`
		} `json:"userIps"`
	} `json:"userIpSegs"`
}
