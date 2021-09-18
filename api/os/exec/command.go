package exec

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os/exec"
)

func CommendWithByte() {
	cmd := exec.Command("echo", "My first command comes from golang")
	fmt.Println(cmd.Args)
	pipe, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Couldn't obtain the studo pipe for commmand No.o: %s\n", err)
		return
	}
	if err := cmd.Start(); err != nil {
		fmt.Printf("Error: %s \n", err)
		return
	}
	output := make([]byte, 10)
	n, err := pipe.Read(output)
	if err != nil {
		fmt.Printf("Error could't read data from the pipe, %s\n", err)
		return
	}
	fmt.Printf("%s\n", output[:n])
}

func ExecCommandWithBuffer() {
	cmd := exec.Command("echo", "My first command comes from golang")
	fmt.Println(cmd.Args)
	pipe, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Couldn't obtain the studo pipe for commmand No.o: %s\n", err)
		return
	}
	if err := cmd.Start(); err != nil {
		fmt.Printf("Error: %s \n", err)
		return
	}
	var outputBufo bytes.Buffer
	for {
		output := make([]byte, 5)
		n, err := pipe.Read(output)
		if err != nil {
			if err == io.EOF {
				fmt.Println("end!")
				break
			} else {
				fmt.Printf("Error could't read data from the pipe, %s\n", err)
				return
			}
		}
		if n > 0 {
			outputBufo.Write(output[:n])
		}
	}
	fmt.Printf("%s\n", outputBufo.String())
}

func ExecCommandWithBufio() {
	cmd := exec.Command("echo", "My first command comes from golang")
	fmt.Println(cmd.Args)
	pipe, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Couldn't obtain the studo pipe for commmand No.o: %s\n", err)
		return
	}
	if err := cmd.Start(); err != nil {
		fmt.Printf("Error: %s \n", err)
		return
	}
	reader := bufio.NewReader(pipe)
	output, _, err := reader.ReadLine()
	if err != nil {
		fmt.Printf("Couldn't obtain the studo pipe for commmand No.o: %s\n", err)
		return
	}
	fmt.Printf("%s\n", string(output))
}

func ExecCommandWithBufio2() {
	cmd := exec.Command("echo", "My first command comes from golang")
	fmt.Println(cmd.Args)
	pipe, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Couldn't obtain the studo pipe for commmand No.o: %s\n", err)
		return
	}
	if err := cmd.Start(); err != nil {
		fmt.Printf("Error: %s \n", err)
		return
	}
	reader := bufio.NewReader(pipe)
	for {
		output, prefix, err := reader.ReadLine()
		if prefix { // 是否是最后一行
			break
		}
		if err == io.EOF {
			fmt.Println("end!")
			break
		}
		fmt.Printf("%s\n", string(output))
	}
}
