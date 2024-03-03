package removenthfromend

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	size := GetLinkedListLength(head)
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

func GetLinkedListLength(head *ListNode) int {
	var tail *ListNode = nil
	tail = head
	it := 0
	for tail.Next != nil {
		tail = tail.Next
		it += 1
	}
	return it + 1
}

func GenerateLinkedList(vals *[]int) *ListNode {
	if len(*vals) == 0 {
		return nil
	}
	head := NewNode((*vals)[0])
	var curr *ListNode = nil
	i := 1
	curr = head
	for ; i < len(*vals); i += 1 {
		curr.Next = NewNode((*vals)[i])
		curr = curr.Next
	}

	return head
}
