package osbasic

import (
	"fmt"
	"os"
)

func DirOperation() {
	file, err := os.Open("D:\\Program Files\\WorkSpace\\go\\src\\GoGoGo")
	if err != nil {
		fmt.Println("err： ", err)
		return
	}

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("err： ", err)
		return
	}
	fmt.Println("name: \t", fileInfo.Name())       // 文件名称
	fmt.Println("size: \t", fileInfo.Size())       // 文件夹大小为 0
	fmt.Println("isDir: \t", fileInfo.IsDir())     // 是否是文件夹
	fmt.Println("modTime: \t", fileInfo.ModTime()) // 最后修改时间
	fmt.Println("mode: \t", fileInfo.Mode())       // 权限
	fmt.Println(file.Readdirnames(-1))             // 查看当前目录下文件、目录名称

	fmt.Println("目录遍历（非递归）")
	file2, err := os.Open("D:\\Program Files\\WorkSpace\\go\\src\\GoGoGo")
	fileInfos, _ := file2.Readdir(-1)
	for _, fileInfo := range fileInfos {
		fmt.Println(fileInfo.Name(), fileInfo.IsDir())
	}

	// 重命名  （也可用于目录的重命名），也可用于文件移动
	// osbasic.Rename("api/osbasic/testwriter2.txt", "api/osbasic/testwriter.txt")

	os.Mkdir("a", os.ModePerm)
	os.MkdirAll("a/b/c/d", os.ModePerm)
	// 删除目录，只可删除空文件夹（可用于文件）
	err = os.Remove("a")
	if err != nil {
		fmt.Println("err: ", err)
	}
	// 递归删除
	os.RemoveAll("a")
}
