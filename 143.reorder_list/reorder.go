package reorder

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReorderList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	var prev *ListNode
	for slow != nil {
		// in
		slow.Next, prev, slow = prev, slow, slow.Next
	}

	first := head
	for prev.Next != nil {
		first.Next, first = prev, first.Next
		prev.Next, prev = first, prev.Next
	}

	return head
}
