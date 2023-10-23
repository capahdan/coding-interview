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

func MergeTwoLists(list1, list2 *ListNode) *ListNode {
	myList := LinkedList{}

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			myList.Append(list1.Val)
			list1 = list1.Next
		} else {
			myList.Append(list2.Val)
			list2 = list2.Next
		}
	}

	for list1 != nil {
		myList.Append(list1.Val)
		list1 = list1.Next
	}

	for list2 != nil {
		myList.Append(list2.Val)
		list2 = list2.Next
	}

	return myList.head
}
