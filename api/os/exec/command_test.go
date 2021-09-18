package exec

import (
	"fmt"
	"testing"
)

func TestCommend(t *testing.T) {
	CommendWithByte()
}

func TestExecCommand(t *testing.T) {
	ExecCommandWithBuffer()
}

func TestExecCommandWithBufio(t *testing.T) {
	ExecCommandWithBufio2()
}

func TestExecCommandWithBufio2(t *testing.T) {
	fmt.Println("缓冲区bufio一次性读取4096个字节，循环读取")
	ExecCommandWithBufio2()
}
