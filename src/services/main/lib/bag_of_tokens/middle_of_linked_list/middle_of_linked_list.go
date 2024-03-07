package middleoflinkedlist

import "project/src/services/main/lib/utils"

type ListNode = utils.ListNode

func MiddleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
