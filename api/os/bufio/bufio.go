package bufio

import (
	"bufio"
	"fmt"
	"os"
)

func Reader() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("err: ",err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	content := make([]byte, 3)
	read, err := reader.Read(content)
	fmt.Println(read, err, content[:read])

	// 读取一个字节
	fmt.Println(reader.ReadByte())
	// 到 | 为止
	fmt.Println(reader.ReadBytes('|'))
	// 到 | 为止
	fmt.Println(reader.ReadString('|'))
	// 读取一行
	fmt.Println(reader.ReadLine())
	fmt.Println(reader.ReadSlice('|'))
}

func Writer() {
	// 相对于本包的路径
	file, err := os.Create("testwriter.txt")
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	// 写入缓存
	fmt.Println(writer.WriteString("aaaaawriter"))
	// 缓存->本地
	writer.Flush()
}
