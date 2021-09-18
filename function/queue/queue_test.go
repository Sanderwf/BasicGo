package queue

import "testing"

func TestQueueNode_Length(t *testing.T) {
	quequ:= new(QueueNode)
	quequ.Create("string",1,[]bool{true,false})
	quequ.Print()
}
