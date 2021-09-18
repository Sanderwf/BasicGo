package NowCoder

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestReverseList(t *testing.T) {
	oldNode := new(ListNode)
	oldNode.Val = 1
	node2 := new(ListNode)
	node2.Val = 2
	node3 := new(ListNode)
	node3.Val = 3
	oldNode.Next = node2
	node2.Next = node3

	marshal, err := json.Marshal(oldNode)
	if err == nil {
		fmt.Println(string(marshal))
	}

	newNode := ReverseList(oldNode)
	marshal, err = json.Marshal(newNode)
	if err == nil {
		fmt.Println(string(marshal))
	}
}
