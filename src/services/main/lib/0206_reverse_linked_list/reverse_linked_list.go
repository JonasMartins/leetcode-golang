package reverselinkedlist

import "project/src/services/main/lib/utils"

type ListNode = utils.ListNode

func ReverseList(head *ListNode) *ListNode {
	var aux *ListNode = nil
	for head != nil {
		var buff *ListNode = head.Next
		head.Next = aux
		aux = head
		head = buff
	}
	return aux
}
