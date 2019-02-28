package main

import "leetcode/kits"

func addTwoNumbers(l1 *kits.ListNodes, l2 *kits.ListNodes) *kits.ListNodes {
	resPre := &kits.ListNodes{}
	cur := resPre
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10

		cur.Next = &kits.ListNodes{Val: sum % 10}
		cur = cur.Next
	}

	return resPre.Next
}
