package reorderlist

import (
	reverselinkedlist "project/src/services/main/lib/0206_reverse_linked_list"
	"project/src/services/main/lib/utils"
)

type ListNode = utils.ListNode

func ReorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	middle := Middle(head)
	reverse := reverselinkedlist.ReverseList(middle)
	Merge(head, reverse)
}

func Middle(list *ListNode) *ListNode {
	var aux *ListNode = nil
	slow := list
	fast := list

	for fast != nil && fast.Next != nil {
		aux = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	aux.Next = nil
	return slow
}

func Merge(list1 *ListNode, list2 *ListNode) {
	for list2 != nil {
		var next *ListNode = nil
		next = list1.Next
		list1.Next = list2
		list1 = list2
		list2 = next
	}
}
