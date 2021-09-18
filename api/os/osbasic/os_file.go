package osbasic

import (
	"fmt"
	"io"
	"os"
)

func FileOperation(){
	file, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return
	}
	defer file.Close()

	// 写完之后光标在最后面，此时直接读取为空
	file.WriteString("123")

	content := make([]byte, 3)
	// 光标在当前位置往前移动两个位置
	file.Seek(-2, 1)

	fmt.Println(file.Read(content))
	fmt.Println(content)

	// 光标在起始位置,12被替换为ad
	file.Seek(0,0)
	file.WriteString("ad")

	file.Seek(-2, 1)
	fmt.Println(file.Read(content))
	fmt.Println(content)
}

func ReadFile(path string)  {
	file, err := os.Open(path)
	if err != nil{
		fmt.Println("err: ", err)
		return
	}
	defer file.Close()

	content := make([]byte, 3)
	for {
		// n: 读取的字节的长度
		n, err := file.Read(content)
		if err != nil {
			if err != io.EOF {
				panic("xxxx")
				fmt.Println(err)
			}
			break
		}
		fmt.Print(string(content[:n]))
	}
}