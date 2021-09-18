package pipe

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"
)
/**
os.pipe两个进程间通信（std.in/std.out）
	命名管道FIFO，以文件的形式存在于文件系统中
channel两个协程间通信
 */

func PipeDeadlock() {
	fmt.Println("io.Pipe会返回一个reader和writer,对reader读取（或写入writer）后，" +
		"进程会被锁住，直到writer有新数据流进入或关闭（或reader把数据读走）。")
	reader, writer := io.Pipe()
	writer.Write([]byte("HELLO"))
	defer writer.Close()

	buffer := make([]byte, 100)
	reader.Read(buffer)
	fmt.Println(string(buffer))
}

func Pipe() {
	reader, writer := io.Pipe()
	lock := make(chan int)
	defer writer.Close()
	// 创建goroutine给reader
	go func() {
		buffer := make([]byte, 100)
		reader.Read(buffer)
		println(string(buffer))
		lock <- 1
	}()
	writer.Write([]byte("hello"))
	<-lock
}

/*
$ cat input
hello world
welcome to world
turtle

$ cat input | grep turtle
turtle
*/
func Pipe_CatGrep() {
	fPtr, _ := os.OpenFile("input", os.O_RDONLY, 0755)
	cmd := exec.Command("grep", "turtle")
	r, w := io.Pipe()
	cmd.Stdin = r
	cmd.Stdout = os.Stdout
	go func() {
		io.Copy(w, fPtr)
		w.Close()
	}()
	cmd.Run()
	time.Sleep(time.Second*5)
}
