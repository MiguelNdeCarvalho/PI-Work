package linkedlist

import (
	"errors"
)

type Node struct {
	Val  interface{}
	next *Node
	prev *Node
}
type List struct {
	head *Node
	tail *Node
}

var ErrEmptyList = errors.New("empty list")

func (e *Node) Next() *Node {
	return e.next
}

func (e *Node) Prev() *Node {
	return e.prev
}

func NewList(args ...interface{}) *List {
	list := &List{}
	for _, item := range args {
		list.PushBack(item)
	}

	return list
}

func (l *List) PushFront(v interface{}) {
	node := &Node{Val: v, next: l.head}
	if l.head == nil && l.tail == nil {
		l.head, l.tail = node, node
		return
	}
	l.head.prev = node
	l.head = node
}

func (l *List) PushBack(v interface{}) {
	node := &Node{Val: v, prev: l.tail}
	if l.head == nil && l.tail == nil {
		l.head, l.tail = node, node
		return
	}
	l.tail.next = node
	l.tail = node
}

func (l *List) PopFront() (interface{}, error) {
	if l.head == nil {
		return nil, ErrEmptyList
	}

	result := l.head.Val
	l.head = l.head.next
	if l.head == nil {
		l.tail = nil
	} else {
		l.head.prev = nil
	}
	return result, nil
}

func (l *List) PopBack() (interface{}, error) {
	if l.tail == nil {
		return nil, ErrEmptyList
	}

	result := l.tail.Val
	l.tail = l.tail.prev
	if l.tail == nil {
		l.head = nil
	} else {
		l.tail.next = nil
	}
	return result, nil
}

func (l *List) Reverse() *List {
	headNode := l.head
	for headNode != nil {
		headNode.next, headNode.prev = headNode.prev, headNode.next
		headNode = headNode.prev
	}
	l.head, l.tail = l.tail, l.head
	return l
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}