package ipLimit

import (
	"fmt"
	"testing"
)

func TestVerifyIp(t *testing.T) {
	ipVerifyList := "192.168.1.0-192.172.3.255"
	ip := "192.170.223.1"
	boolen := VerifyIp(ipVerifyList, ip)
	t.Log("是否包含：", boolen)
}

func TestIsIp(t *testing.T) {
	isIp := IsIp("255.255.255.0/12")
	fmt.Println(isIp)
}

func TestIsIp2(t *testing.T) {
	isIp := IsIp2("192.168.0.1/12")
	fmt.Println(isIp)
}