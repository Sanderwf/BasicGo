package osbasic

import (
	"fmt"
	"os"
)

func SomeOperation()  {
	dir, err := os.Getwd()
	fmt.Println(dir)
	fmt.Println(err)

	// 获取环境变量
	GOROOT := os.Getenv("GOROOT")
	PATH := os.Getenv("PATH")
	fmt.Println(GOROOT, PATH)

}
