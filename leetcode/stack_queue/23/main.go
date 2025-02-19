package main

import (
	"container/heap"
	"fmt"
)

// 테스트를 위한 헬퍼 함수
func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Println("nil")
}

func main() {
	// 예시: 두 개의 리스트를 병합
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	list3 := &ListNode{Val: 2, Next: &ListNode{Val: 6}}

	lists := []*ListNode{list1, list2, list3}
	merged := mergeKLists(lists)
	printList(merged)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNodeHeap []*ListNode

func (h ListNodeHeap) Len() int {
	return len(h)
}
func (h ListNodeHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h ListNodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *ListNodeHeap) Push(x any) {
	*h = append(*h, x.(*ListNode))
}

func (h *ListNodeHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func mergeKLists(lists []*ListNode) *ListNode {
	var h ListNodeHeap
	heap.Init(&h)

	for _, list := range lists {
		if list != nil {
			heap.Push(&h, list)
		}
	}
	dummy := &ListNode{Val: 0}
	tail := dummy

	for h.Len() > 0 {
		minNode := heap.Pop(&h).(*ListNode)
		tail.Next = minNode
		tail = tail.Next
		if minNode.Next != nil {
			heap.Push(&h, minNode.Next)
		}
	}
	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
