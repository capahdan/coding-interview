package valid_parentheses

type ListNode struct {
	Val  int
	Next *ListNode
}

func HasCycle(head *ListNode) bool {

	// we create 2 pointer , 1 slow and 1 move 2 times
	if head == nil {
		return false
	}

	slowPointer := head
	fastPointer := head

	for slowPointer.Next != nil && fastPointer.Next != nil {
		slowPointer = slowPointer.Next
		fastPointer = fastPointer.Next.Next

		if slowPointer == fastPointer {
			return true
		}
	}

	return false
}
