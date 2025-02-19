package main

import "container/list"

func main() {

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Deque는 container/list를 기반으로 한 연결리스트 기반 Deque 구현체다.
type Deque struct {
	data *list.List
}

// NewDeque는 빈 Deque를 생성한다.
func NewDeque() *Deque {
	return &Deque{data: list.New()}
}

// AddLast는 Deque의 뒤쪽에 값을 추가한다.
func (d *Deque) AddLast(val int) {
	d.data.PushBack(val)
}

// PollFirst는 Deque의 앞쪽 요소를 반환하고 제거한다.
// Deque가 비어있으면 (0, false)를 반환한다.
func (d *Deque) PollFirst() (int, bool) {
	front := d.data.Front()
	if front == nil {
		return 0, false
	}
	val := front.Value.(int)
	d.data.Remove(front)
	return val, true
}

// PollLast는 Deque의 뒤쪽 요소를 반환하고 제거한다.
// Deque가 비어있으면 (0, false)를 반환한다.
func (d *Deque) PollLast() (int, bool) {
	back := d.data.Back()
	if back == nil {
		return 0, false
	}
	val := back.Value.(int)
	d.data.Remove(back)
	return val, true
}

// Size는 Deque에 저장된 요소의 개수를 반환한다.
func (d *Deque) Size() int {
	return d.data.Len()
}

// isPalindrome은 연결리스트를 Deque에 저장한 후, 양쪽 끝에서 값을 꺼내어 회문 여부를 판단한다.
func isPalindrome(head *ListNode) bool {
	d := NewDeque()
	// 연결리스트를 순회하며 모든 값을 Deque에 저장한다.
	for node := head; node != nil; node = node.Next {
		d.AddLast(node.Val)
	}

	// Deque의 양쪽 끝에서 하나씩 꺼내어 비교한다.
	for d.Size() > 1 {
		first, _ := d.PollFirst()
		last, _ := d.PollLast()
		if first != last {
			return false
		}
	}
	return true
}
