package NowCoder

/*
输入一个链表，反转链表后，输出新链表的表头。
输入：{1,2,3}
返回值：{3,2,1}
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(oldNode *ListNode) *ListNode {
	if oldNode == nil || oldNode.Next == nil {
		return oldNode
	}
	var newNode *ListNode
	for oldNode != nil {
		oldNextNode := oldNode.Next
		oldNode.Next = newNode
		newNode = oldNode
		oldNode = oldNextNode
	}
	return newNode
}
