package add_two_numbers

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		sum, endOfTen int
		result        = &ListNode{}
		current       = result
	)

	for l1 != nil || l2 != nil || endOfTen > 0 {
		sum = endOfTen

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		if sum >= 10 {
			endOfTen = sum / 10
			sum %= 10
		} else {
			endOfTen = 0
		}

		current.Next = &ListNode{Val: sum}
		current = current.Next
	}

	return result.Next
}
