package circular

import (
	"fmt"
)

//Queue a  simple implementation of a queue structure, usign a byte slice
type Queue struct {
	size int
	data []byte
}

func newQueue(size int) *Queue {
	return &Queue{size: size}
}

func (q *Queue) push(b byte) error {
	if len(q.data) < q.size {
		q.data = append(q.data, b)
		return nil
	}
	return fmt.Errorf("Failed to push value : %v", b)
}

func (q *Queue) forcePush(b byte) error {
	if len(q.data) == q.size {
		q.pop()
	}
	return q.push(b)
}

func (q *Queue) pop() (result byte, err error) {
	if len(q.data) > 0 {
		result := q.data[0]

		if len(q.data) > 1 {
			q.data = q.data[1:]
		} else {
			q.data = []byte{}
		}

		return result, err
	}
	return result, fmt.Errorf("Failed to pop value from empty queue")
}

func (q *Queue) clear() {
	q.data = []byte{}
}

type Buffer struct {
	data Queue
}

// NewBuffer returns a new buffer with buffer size.... don't pass in a negative size silly
func NewBuffer(size int) *Buffer {
	buffer := Buffer{}
	data := []byte{}
	buffer.data = Queue{size: size, data: data}
	return &buffer
}

// ReadByte reads and return the oldest byte in the buffer
func (b *Buffer) ReadByte() (byte, error) {
	return b.data.pop()
}

// WriteByte writes the given byte into the buffer on an empty slot, returns error if buffer
// already full
func (b *Buffer) WriteByte(c byte) error {
	return b.data.push(c)
}

// Overwrite writes the given byte into the buffer on an empty slot if available, or overides
// the oldest value in the buffer. Doesn't return any errors
func (b *Buffer) Overwrite(c byte) {
	b.data.forcePush(c)
}

// Reset resets the buffer, probably clearing all the bytes in the buffer
func (b *Buffer) Reset() {
	b.data.clear()
}