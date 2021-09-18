package ipLimit

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Ip2Int(ip string) int64 {
	if len(ip) == 0 {
		return 0
	}
	bits := strings.Split(ip, ".")
	if len(bits) < 4 {
		return 0
	}
	b0 := String2Int(bits[0])
	b1 := String2Int(bits[1])
	b2 := String2Int(bits[2])
	b3 := String2Int(bits[3])

	var sum int64
	sum += int64(b0) << 24
	sum += int64(b1) << 16
	sum += int64(b2) << 8
	sum += int64(b3)
	fmt.Println("ip: ", ip, ", sum: ", sum)
	return sum
}

func String2Int(in string) (out int) {
	out, _ = strconv.Atoi(in)
	return
}

func VerifyIp(ipVerifyList, ip string) bool {
	ipSlice := strings.Split(ipVerifyList, `-`)
	if len(ipSlice) < 0 {
		return false
	}
	if Ip2Int(ip) >= Ip2Int(ipSlice[0]) && Ip2Int(ip) <= Ip2Int(ipSlice[1]) {
		return true
	}
	return false
}

func IsIp(ip string)bool {
	if m, _ := regexp.MatchString("^(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)$", ip); m {
		return true
	}
	return false
}

func IsIp2(ip string)bool {
	if m, _ := regexp.MatchString("\"(?:(?:25[0-5]|2[0-4]/d|[01]?/d?/d)/.){3}(?:25[0-5]|2[0-4]/d|[01]?/d?/d)\"", ip); m {
		return true
	}
	return false
}

