package linkedlist

import (
	"errors"
)

type Element struct {
	data int
 	next *Element
}

type List struct {
 	head *Element
 	size int
}

func New(l []int) *List {
	list := List{head: nil, size: 0}
	for _, v := range l {
		list.Push(v)
	}
	return &list
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Push(data int) {
	newElement := &Element{data: data, next: nil}
	defer func() { l.size++ }()
	if l.head == nil {
		l.head = newElement
		return
	}

	e := l.head
	for e.next != nil {
		e = e.next
	}
	e.next = newElement
}

func (l *List) Pop() (int, error) {
	if l.head == nil {
		return 0, errors.New("no head")
	}
	if l.head.next == nil {
		d := l.head.data
		l.head = nil
		l.size--
		return d, nil
	}
	prev := l.head
	e := l.head
	for e.next != nil {
		prev, e = e, e.next
	}
	prev.next = nil
	l.size--
	return e.data, nil
}

func (l *List) Array() []int {
	var ret []int
	for {
		d, err := l.Pop()
		if err != nil {
			break
		}
		ret = append([]int{d}, ret...)
	}
	return ret
}

func (l *List) Reverse() *List {
	rev := New(nil)
	for {
		d, err := l.Pop()
		if err != nil {
			break
		}
		rev.Push(d)
	}
	return rev
}