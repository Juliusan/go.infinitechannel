package queue

// PriorityQueue is an improvement of SimpleQueue, which can prioritise elements.
type PriorityQueue struct {
	buf               []interface{}
	head, tail, count int
	priorityFun       func(interface{}) bool
}

var _ Queue = &PriorityQueue{}

// New constructs and returns a new Queue. Code duplication needed for benchmarks.
func NewDefaultPriorityQueue() *PriorityQueue {
	return NewPriorityQueue(func(interface{}) bool { return false })
}

func NewPriorityQueue(fun func(interface{}) bool) *PriorityQueue {
	return &PriorityQueue{
		buf:         make([]interface{}, minQueueLen),
		priorityFun: fun,
	}
}

// Length returns the number of elements currently stored in the queue.
func (q *PriorityQueue) Length() int {
	return q.count
}

// resizes the queue to fit exactly twice its current contents
// this can result in shrinking if the queue is less than half-full
func (q *PriorityQueue) resize() {
	newBuf := make([]interface{}, q.count<<1)

	if q.tail > q.head {
		copy(newBuf, q.buf[q.head:q.tail])
	} else {
		n := copy(newBuf, q.buf[q.head:])
		copy(newBuf[n:], q.buf[:q.tail])
	}

	q.head = 0
	q.tail = q.count
	q.buf = newBuf
}

// Add puts an element to the start of end of the queue, depending
// on the result of priorityFun.
func (q *PriorityQueue) Add(elem interface{}) {
	if q.count == len(q.buf) {
		q.resize()
	}

	priority := q.priorityFun(elem)
	if priority {
		q.head = (q.head - 1) & (len(q.buf) - 1)
		q.buf[q.head] = elem
	} else {
		q.buf[q.tail] = elem
		// bitwise modulus
		q.tail = (q.tail + 1) & (len(q.buf) - 1)
	}
	q.count++
}

// Peek returns the element at the head of the queue. This call panics
// if the queue is empty.
func (q *PriorityQueue) Peek() interface{} {
	if q.count <= 0 {
		panic("queue: Peek() called on empty queue")
	}
	return q.buf[q.head]
}

// Get returns the element at index i in the queue. If the index is
// invalid, the call will panic. This method accepts both positive and
// negative index values. Index 0 refers to the first element, and
// index -1 refers to the last.
func (q *PriorityQueue) Get(i int) interface{} {
	// If indexing backwards, convert to positive index.
	if i < 0 {
		i += q.count
	}
	if i < 0 || i >= q.count {
		panic("queue: Get() called with index out of range")
	}
	// bitwise modulus
	return q.buf[(q.head+i)&(len(q.buf)-1)]
}

// Remove removes and returns the element from the front of the queue. If the
// queue is empty, the call will panic.
func (q *PriorityQueue) Remove() interface{} {
	if q.count <= 0 {
		panic("queue: Remove() called on empty queue")
	}
	ret := q.buf[q.head]
	q.buf[q.head] = nil
	// bitwise modulus
	q.head = (q.head + 1) & (len(q.buf) - 1)
	q.count--
	// Resize down if buffer 1/4 full.
	if len(q.buf) > minQueueLen && (q.count<<2) == len(q.buf) {
		q.resize()
	}
	return ret
}
