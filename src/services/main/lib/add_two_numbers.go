package lib

import (
	"fmt"
	"math/big"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1 := GetSumFromLinkedList(l1)
	s2 := GetSumFromLinkedList(l2)
	sum := big.NewInt(0)
	sum.Add(&s1, &s2)
	return GetReversedLinkedListFromNumber(sum)
}
func PrintLinkedList(l *ListNode) {
	cursor := l
	for {
		fmt.Printf("%d,", cursor.Val)
		if cursor.Next == nil {
			break
		}
		cursor = cursor.Next
	}
	fmt.Print("\n")
}

func GetSliceFromNumber(n *big.Int) *[]uint8 {
	slice := []uint8{}
	var buff *big.Int = big.NewInt(0)
	var aux *big.Int
	zero := big.NewInt(0)
	remainder := big.NewInt(0)
	*buff = *n
	for {
		if buff.Cmp(zero) == 0 {
			break
		}
		aux = GetRemainder(buff, remainder)
		slice = append(slice, uint8(aux.Uint64()))
	}
	if len(slice) == 0 {
		slice = append(slice, 0)
	}
	return &slice
}

func GetRemainder(n *big.Int, r *big.Int) *big.Int {
	n.QuoRem(n, big.NewInt(10), r)
	return r
}

func GetReversedLinkedListFromNumber(n *big.Int) *ListNode {
	slice := GetSliceFromNumber(n)
	result := BuildLinkedListFromArray(slice)
	return result
}

func GetSumFromLinkedList(l *ListNode) big.Int {
	tail := l
	len := uint8(1)
	var sum big.Int
	var base big.Int
	var i big.Int
	var iSum big.Int
	var buff *big.Int
	base.SetUint64(10.0)
	sum.SetUint64(0)
	i.SetUint64(0)
	iSum.SetUint64(1)
	for {
		if tail.Next == nil {
			break
		}
		tail = tail.Next
		len += 1
	}
	tail = l
	len -= 1
	for {
		buff = new(big.Int).Exp(&base, &i, nil)
		sum.Add(&sum, buff.Mul(buff, big.NewInt(int64(tail.Val))))
		// sum += float64(tail.Val) * math.Pow(10.0, float64(i))
		if len == 0 {
			break
		}
		len -= 1
		tail = tail.Next
		i.Add(&i, &iSum)
	}

	return sum
}

func BuildLinkedListFromArray(values *[]uint8) *ListNode {
	head := ListNode{
		Val:  0,
		Next: nil,
	}
	head.Val = int((*values)[0])
	i := 1
	cursor := &head
	for ; i < len((*values)); i += 1 {
		aux := ListNode{
			Val:  int((*values)[i]),
			Next: nil,
		}
		cursor.Next = &aux
		cursor = &aux
	}
	return &head
}
