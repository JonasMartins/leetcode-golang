package palindromelinkedlist

import (
	reverselinkedlist "project/src/services/main/lib/0206_reverse_linked_list"
	"project/src/services/main/lib/utils"
)

type ListNode = utils.ListNode

func IsPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil {
		slow = slow.Next
	}
	slow = reverselinkedlist.ReverseList(slow)
	isPalindrome := true
	for slow != nil {
		if slow.Val != head.Val {
			isPalindrome = false
			break
		}
		slow = slow.Next
		head = head.Next
	}
	return isPalindrome
}
