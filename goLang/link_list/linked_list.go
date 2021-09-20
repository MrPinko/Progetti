package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	l3 := &ListNode{Val: 2}
	l2 := &ListNode{Val: 4, Next: l3}
	l1 := &ListNode{Val: 6, Next: l2}
	List := &ListNode{Next: l1}

	List.Display()

}

func (l ListNode) Display() {
	for l.Next != nil {
		fmt.Printf("%v -> ", l.Next.Val)
		l.Next = l.Next.Next
	}
	fmt.Println()
}
