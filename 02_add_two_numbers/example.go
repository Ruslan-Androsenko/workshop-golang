package add_two_numbers

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	iter := head
	carry := 0
	for l1 != nil || l2 != nil || carry > 0 {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		sum += carry
		iter.Next = &ListNode{Val: sum % 10}
		carry = sum / 10
		iter = iter.Next
	}
	return head.Next
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	var ret *ListNode
	var perNode *ListNode
	s := 0
	i := 0
	for {

		if l1 != nil {
			s += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			s += l2.Val
			l2 = l2.Next
		}

		v := s % 10
		s /= 10
		newNode := ListNode{v, nil}
		if i == 0 {
			ret = &newNode

		} else {
			perNode.Next = &newNode

		}
		perNode = &newNode
		if l1 == nil && l2 == nil && s == 0 {
			break
		}
		i++
	}
	return ret
}
