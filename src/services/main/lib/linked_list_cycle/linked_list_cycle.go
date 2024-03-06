package linkedlistcycle

import (
	"project/src/services/main/lib/utils"
)

func HasCycle(head *utils.ListNode) bool {
	result := false
	if head == nil || head.Next == nil {
		return result
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			result = true
			break
		}
	}
	return result
}
