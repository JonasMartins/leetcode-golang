package removenthfromend

import "project/src/services/main/lib/utils"

type ListNode = utils.ListNode

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	size := utils.GetLinkedListLength(head)
	nth := size - n
	nthEl := GetLinkedListNthElement(head, nth)
	if head == nthEl {
		head = head.Next
		return head
	}
	var curr *ListNode = nil
	curr = head

	for curr.Next != nil {
		if curr.Next == nthEl {
			curr.Next = curr.Next.Next
			break
		}
		curr = curr.Next
	}

	return head
}

func NewNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

func GetLinkedListNthElement(head *ListNode, nth int) *ListNode {
	var el *ListNode = nil
	el = head
	it := 1
	for el.Next != nil && it <= nth {
		el = el.Next
		it += 1
	}
	return el
}
