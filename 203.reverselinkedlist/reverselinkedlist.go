package reverselinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var revHead *ListNode

	for head != nil {
		tmp := head.Next
		head.Next = revHead
		revHead = head
		head = tmp
	}
	return revHead
}
