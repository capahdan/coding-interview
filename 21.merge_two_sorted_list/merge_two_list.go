package valid_parentheses

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	head *ListNode
}

func (ll *LinkedList) Append(data int) {
	newNode := &ListNode{Val: data, Next: nil}

	if ll.head == nil {
		ll.head = newNode
	} else {
		current := ll.head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
}

func MergeTwoLists(l1, l2 *ListNode) *ListNode {
	// Create a dummy node to start the merged list
	dummy := &ListNode{}
	// Pointer to the current node in the merged list
	current := dummy

	// Iterate until either of the lists becomes nil
	for l1 != nil && l2 != nil {
		// Compare values of the two lists and append the smaller one to the merged list
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		// Move the current pointer forward
		current = current.Next
	}

	// Append the remaining nodes of the non-nil list
	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}

	// Return the merged list starting from the next of the dummy node
	return dummy.Next
}
