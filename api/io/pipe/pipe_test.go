package pipe

import "testing"

func TestPipeDemo(t *testing.T) {
	PipeDeadlock()
}

func TestPipe(t *testing.T) {
	Pipe()
}

func TestPipe_catGrep(t *testing.T) {
	Pipe_CatGrep()
}
