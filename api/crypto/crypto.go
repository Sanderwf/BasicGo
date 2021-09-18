package crypto

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func CryptoDemo() {
	// 不同的算法返回的字节长度不一样
	text := "我是王飞"
	fmt.Println("\n", GetMd5Str(text))
	fmt.Println("\n", GetMd5Str2())
	fmt.Println("\n", GetSha1Str(text))
	fmt.Println("\n", GetSha256Str(text))
	fmt.Println("\n", GetSha512Str(text))
}

func GetMd5Str(str string) string {
	return fmt.Sprintf("%X", md5.Sum([]byte(str)))
}

func GetMd5Str2() string {
	fmt.Println("分批写入得到的结果是一样的")
	md5hasher := md5.New()
	md5hasher.Write([]byte("我是"))
	md5hasher.Write([]byte("王飞"))
	return fmt.Sprintf("%X", md5hasher.Sum(nil))
}

func GetSha1Str(str string) string {
	return fmt.Sprintf("%X", sha1.Sum([]byte(str)))
}


func GetSha256Str(str string) string {
	return fmt.Sprintf("%X", sha256.Sum256([]byte(str)))
}

func GetSha512Str(str string) string {
	return fmt.Sprintf("%x", sha512.Sum512([]byte(str)))
}
