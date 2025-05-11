package data_structs

type List[T any] struct {
	head *ListNode[T]
	size int
}

type ListNode[T any] struct {
	Val  T
	Next *ListNode[T]
}

func NewList[T any](head *ListNode[T]) List[T] {
	if head == nil {
		panic("NewList: head should be not nil value")
	}

	list := List[T]{head: head}
	list.setSize()

	return list
}

func (this *List[T]) setSize() {
	size := 1
	next := this.head
	for next.Next != nil {
		size++
		next = next.Next
	}
	this.size = size
}

func (this *List[T]) Size() int {
	return this.size
}

func (this *List[T]) Head() *ListNode[T] {
	return this.head
}

func (this *List[T]) Tail() *ListNode[T] {
	next := this.head
	for next.Next != nil {
		next = next.Next
	}

	return next
}

func (this *List[T]) PreTail() *ListNode[T] {
	preTail := new(ListNode[T])
	next := this.head
	for next.Next != nil {
		if next.Next == this.Tail() {
			return next
		}
		next = next.Next
	}

	return preTail
}

func (this *List[T]) ToSlice() []T {
	s := make([]T, 0, this.Size())
	next := this.head
	for next.Next != nil {
		s = append(s, next.Val)
		next = next.Next
	}
	s = append(s, next.Val)

	return s
}

func (this *List[T]) Push(value T) int {
	tail := this.Tail()
	tail.Next = &ListNode[T]{Val: value, Next: nil}
	this.size++

	return this.Size()
}

func (this *List[T]) Pop() *ListNode[T] {
	if this.Size() == 1 {
		return nil
	}

	preTail := new(ListNode[T])
	tail := this.Tail()

	next := this.head
	for next.Next != nil {
		if next.Next == tail {
			preTail = next
		}
		next = next.Next
	}
	preTail.Next = nil
	this.size--

	return tail
}

func (this *List[T]) TailToHead() {
	if this.Size() == 1 {
		return
	}
	tail := this.Pop()
	tail.Next = this.head
	this.head = tail
	this.size++
}
