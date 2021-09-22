package queue

// LimitedPriorityQueuePriorityQueue is an improvement of SimpleQueue, which
// can prioritise elements and limit its growth. When the limit is reached, oldest
// elements are discarded.
type LimitedPriorityQueue struct {
	buf         []interface{}
	head        int
	ptail       int
	tail        int
	count       int
	priorityFun func(interface{}) bool
	limit       int
}

var _ Queue = &LimitedPriorityQueue{}

// New constructs and returns a new Queue. Code duplication needed for benchmarks.
func NewDefaultLimitedPriorityQueue() *LimitedPriorityQueue {
	return NewLimitLimitedPriorityQueue(Infinity)
}

func NewPriorityLimitedPriorityQueue(fun func(interface{}) bool) *LimitedPriorityQueue {
	return NewLimitedPriorityQueue(fun, Infinity)
}

func NewLimitLimitedPriorityQueue(limit int) *LimitedPriorityQueue {
	return NewLimitedPriorityQueue(func(interface{}) bool { return false }, limit)
}

func NewLimitedPriorityQueue(fun func(interface{}) bool, limit int) *LimitedPriorityQueue {
	return &LimitedPriorityQueue{
		buf:         make([]interface{}, minQueueLen),
		priorityFun: fun,
		limit:       limit,
	}
}

// Length returns the number of elements currently stored in the queue.
func (q *LimitedPriorityQueue) Length() int {
	return q.count
}

// resizes the queue to fit exactly twice its current contents
// this can result in shrinking if the queue is less than half-full
func (q *LimitedPriorityQueue) resize() bool {
	var newSize int
	if q.count >= q.limit {
		return false
	} else {
		newSize = q.count << 1
		if newSize > q.limit {
			newSize = q.limit
		}
	}
	newBuf := make([]interface{}, newSize)

	if q.tail > q.head {
		copy(newBuf, q.buf[q.head:q.tail])
	} else {
		n := copy(newBuf, q.buf[q.head:])
		copy(newBuf[n:], q.buf[:q.tail])
	}

	q.ptail = (q.ptail - q.head) % newSize
	q.head = 0
	q.tail = q.count
	q.buf = newBuf
	return true
}

// Add puts an element to the start of end of the queue, depending
// on the result of priorityFun.
func (q *LimitedPriorityQueue) Add(elem interface{}) bool {
	limitReached := false
	if q.count == len(q.buf) {
		limitReached = !q.resize()
	}
	priority := q.priorityFun(elem)
	if limitReached && !priority && (q.ptail == q.tail) {
		//Not possible to add not priority element in queue full of priority elements
		return false
	}
	if limitReached {
		if q.ptail > q.head {
			copy(q.buf[q.head+1:q.ptail+1], q.buf[q.head:q.ptail])
		} else {
			oldHead := q.buf[q.head]
			if q.ptail > 0 {
				copy(q.buf[1:q.ptail+1], q.buf[:q.ptail])
			}
			lastIndex := len(q.buf) - 1
			q.buf[0] = q.buf[lastIndex]
			if q.head < lastIndex {
				copy(q.buf[q.head+1:], q.buf[q.head:lastIndex])
			}
			q.buf[(q.head+1)%len(q.buf)] = oldHead
		}
		q.head = (q.head + 1) % len(q.buf)
		q.ptail = (q.ptail + 1) % len(q.buf)
	}
	if priority {
		q.head = (q.head - 1) % len(q.buf)
		q.buf[q.head] = elem
		q.ptail = (q.ptail + 1) % len(q.buf)
	} else {
		q.buf[q.tail] = elem
		// bitwise modulus
		q.tail = (q.tail + 1) % len(q.buf)
	}
	q.count++
	return true
}

// Peek returns the element at the head of the queue. This call panics
// if the queue is empty.
func (q *LimitedPriorityQueue) Peek() interface{} {
	if q.count <= 0 {
		panic("queue: Peek() called on empty queue")
	}
	return q.buf[q.head]
}

// Get returns the element at index i in the queue. If the index is
// invalid, the call will panic. This method accepts both positive and
// negative index values. Index 0 refers to the first element, and
// index -1 refers to the last.
func (q *LimitedPriorityQueue) Get(i int) interface{} {
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
func (q *LimitedPriorityQueue) Remove() interface{} {
	if q.count <= 0 {
		panic("queue: Remove() called on empty queue")
	}
	ret := q.buf[q.head]
	q.buf[q.head] = nil
	if q.head == q.ptail {
		q.head = (q.head + 1) % len(q.buf)
	}
	q.head = (q.head + 1) % len(q.buf)
	q.count--
	// Resize down if buffer 1/4 full.
	if len(q.buf) > minQueueLen && (q.count<<2) == len(q.buf) {
		q.resize()
	}
	return ret
}
