package queue

import "fmt"

type QueueNode struct {
	Data interface{}
	Next *QueueNode
}

func (queue *QueueNode) Create(Data ...interface{}) {
	if queue == nil {
		return
	}
	if len(Data) == 0 {
		return
	}

	for _, v := range Data {
		newNode := new(QueueNode)
		newNode.Data = v
		queue.Next = newNode
		queue = queue.Next
	}
}

func (queue *QueueNode) Print() {
	if queue == nil {
		return
	}
	for queue != nil {
		if queue.Data != nil {
			fmt.Print(queue.Data, " ")
		}
		queue = queue.Next
	}
	fmt.Println()
}

func (queue *QueueNode) Length() int {
	if queue == nil {
		return -1
	}
	i := 0
	for queue.Next != nil {
		i++
		queue = queue.Next
	}
	return i
}

// 入列
func (queue *QueueNode) Push(Data interface{}) {
	if queue == nil {
		return
	}
	if Data == nil {
		return
	}
	//找到队列末尾
	for queue.Next != nil {
		queue = queue.Next
	}
	newNode := new(QueueNode)
	newNode.Data = Data
	queue.Next = newNode
}

func (queue *QueueNode) Pop() {
	//队头出列
	if queue == nil {
		return
	}
	//记录列队第一个的节点
	//node:=queue.Next
	//queue.Next=node.Next
	queue.Next = queue.Next.Next
}
